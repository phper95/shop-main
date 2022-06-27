package wechat

import (
	"gitee.com/phper95/pkg/cache"
	"github.com/silenceper/wechat/v2"
	wechatConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"shop/pkg/global"
)

func InitWechat() {
	wc := wechat.NewWechat()
	//这里使用redis保存access_token，也可选择memcache或者自定cache，或者内存
	wc.SetCache(cache.GetRedisClient(cache.DefaultRedisClient))
	cfg := &wechatConfig.Config{
		AppID:          global.CONFIG.Wechat.AppID,
		AppSecret:      global.CONFIG.Wechat.AppSecret,
		Token:          global.CONFIG.Wechat.Token,
		EncodingAESKey: global.CONFIG.Wechat.EncodingAESKey,
	}

	officialAccount := wc.GetOfficialAccount(cfg)

	global.WechatOfficial = officialAccount
}
