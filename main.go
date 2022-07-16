package main

import (
	"context"
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/db"
	"gitee.com/phper95/pkg/mq"
	"gitee.com/phper95/pkg/shutdown"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
	"shop/internal/listen"
	"shop/pkg/base"
	"shop/pkg/casbin"
	"shop/pkg/global"
	"shop/pkg/jwt"
	"shop/pkg/logging"
	"shop/pkg/wechat"
	"shop/routers"
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

	//优雅关闭
	shutdown.NewHook().Close(
		//关闭http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logging.Error("http server shutdown err", err)
			}
		},

		func() {
			//关闭kafka producer（特别是异步生产者，强制关闭会导致丢消息）
			if err := mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Close(); err != nil {
				logging.Error("kafka shutdown err", err)
			}
		},
		func() {
			//关闭mysql
			if err := db.CloseMysqlClient(db.DefaultClient); err != nil {
				logging.Error("mysql shutdown err", err)
			}
		},
		func() {
			//关闭redis
			if err := cache.GetRedisClient(cache.DefaultRedisClient).Close(); err != nil {
				logging.Error("redis shutdown err", err)
			}
		},
	)
	//也可以自己实现优雅关闭
	//signals := make(chan os.Signal, 0)
	//signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//s := <-signals
	//global.LOG.Warn("shop receive system signal:", s)
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//err := server.Shutdown(ctx)
	//if err != nil {
	//	global.LOG.Error("http server error", err)
	//}
	//mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Close()
}
