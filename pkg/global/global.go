package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"shop/conf"
)

var (
	Db             *gorm.DB
	LOG            *zap.SugaredLogger
	CONFIG         conf.Config
	WechatOfficial *officialaccount.OfficialAccount
)

// 加载配置，失败直接panic
func LoadConfig() {
	viper := viper.New()
	//1.设置配置文件路径
	viper.SetConfigFile("conf/config.yml")
	//2.配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//3.将配置映射成结构体
	if err := viper.Unmarshal(&CONFIG); err != nil {
		panic(err)
	}

	//4. 监听配置文件变动,重新解析配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name)
		if err := viper.Unmarshal(&CONFIG); err != nil {
			panic(err)
		}
	})
}
