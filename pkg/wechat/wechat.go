package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	wechatConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"shop/pkg/global"
)

func InitWechat() {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	redisOpts := &cache.RedisOpts{
		Host:        global.CONFIG.Redis.Host,
		Password:    global.CONFIG.Redis.Password,
		Database:    0,
		MaxActive:   global.CONFIG.Redis.MaxActive,
		MaxIdle:     global.CONFIG.Redis.MaxIdle,
		IdleTimeout: 200,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	cfg := &wechatConfig.Config{
		AppID:          global.CONFIG.Wechat.AppID,
		AppSecret:      global.CONFIG.Wechat.AppSecret,
		Token:          global.CONFIG.Wechat.Token,
		EncodingAESKey: global.CONFIG.Wechat.EncodingAESKey,
	}

	officialAccount := wc.GetOfficialAccount(cfg)

	global.WechatOfficial = officialAccount
}
