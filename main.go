package main

import (
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/db"
	"gitee.com/phper95/pkg/mq"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"log"
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
		panic(err)
	}

	err = db.InitMysqlClient(db.DefaultClient, global.CONFIG.Database.User,
		global.CONFIG.Database.Password, global.CONFIG.Database.Host, global.CONFIG.Database.Name)
	if err != nil {
		panic(err)
	}
	global.Db = db.GetMysqlClient(db.DefaultClient).DB

	casbin.InitCasbin(global.Db)

	jwt.Init()

	listen.Init()

	wechat.InitWechat()

	err = mq.InitAsyncKafkaProducer(mq.DefaultKafkaSyncProducer, global.CONFIG.Kafka.Hosts, nil)
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

	global.LOG.Info("[info] start http server listening %s", endPoint)
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()

}
