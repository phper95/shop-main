package constant

import "time"

const (
	ContextKeyUserObj  = "authedUserObj"
	RedisPrefixAuth    = "auth:"
	CASBIN             = "gin-shop"
	WeChatMenu         = "wechat_menus"
	AppRedisPrefixAuth = "app_auth:"
	AppAuthUser        = "app_auth_user:"
	SmsCode            = "sms_code:"
	SmsLength          = 6
	CityList           = "shop-city:"
	OrderInfo          = "order-info:"
	//Header 中传递的参数字段，其携带的值为接口的签名
	HeaderAuthField = "Authorization"

	//Header 中传递的参数字段，其携带的值为发起请求的时间，用于签名失效验证
	HeaderAuthDateField = "Authorization-Date"

	//签名失效时间
	AuthorizationExpire = time.Minute * 30

	RedisKeyPrefixSignature       = "sign:"
	RedisSignatureCacheSeconds    = 300 * time.Second
	HeaderSignTokenTimeoutSeconds = 300 * time.Second
)
