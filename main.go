package main

import (
	"context"
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/db"
	"gitee.com/phper95/pkg/mq"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/listen"
	"shop/pkg/base"
	"shop/pkg/casbin"
	"shop/pkg/global"
	"shop/pkg/jwt"
	"shop/pkg/logging"
	"shop/pkg/wechat"
	"shop/routers"
	"syscall"
	"time"
)

func init() {
	global.LoadConfig()

	global.LOG = base.SetupLogger()

	logging.Init()

	err := cache.InitRedis(cache.DefaultRedisClient, &redis.Options{
		Addr:        global.CONFIG.Redis.Host,
		Password:    global.CONFIG.Redis.Password,
		IdleTimeout: global.CONFIG.Redis.IdleTimeout,
	}, nil)
	if err != nil {
		if err != nil {
			global.LOG.Error("InitRedis error ", err, "client", cache.DefaultRedisClient)
			panic(err)
		}
		panic(err)
	}

	err = db.InitMysqlClient(db.DefaultClient, global.CONFIG.Database.User,
		global.CONFIG.Database.Password, global.CONFIG.Database.Host,
		global.CONFIG.Database.Name)
	if err != nil {
		global.LOG.Error("InitMysqlClient error ", err, "client", db.DefaultClient)
		panic(err)
	}
	global.Db = db.GetMysqlClient(db.DefaultClient).DB

	casbin.InitCasbin(global.Db)

	jwt.Init()

	listen.Init()

	wechat.InitWechat()

	err = mq.InitSyncKafkaProducer(mq.DefaultKafkaSyncProducer, global.CONFIG.Kafka.Hosts, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(global.CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logging.Error("start http server error", err)
		} else {
			fmt.Println("start http server listening", endPoint)
		}
	}()

	signals := make(chan os.Signal, 0)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-signals
	global.LOG.Warn("shop receive system signal:", s)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		global.LOG.Error("http server error", err)
	}
	mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Close()
}
