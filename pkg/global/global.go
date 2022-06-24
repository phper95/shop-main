package global

import (
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"shop/conf"
)

var (
	Db             *gorm.DB
	Shop_VP        *viper.Viper
	LOG            *zap.SugaredLogger
	CONFIG         conf.Config
	WechatOfficial *officialaccount.OfficialAccount
)
