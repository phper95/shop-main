package models

import (
	"gorm.io/datatypes"
	"shop/pkg/constant"
	"shop/pkg/global"
	"time"
)

type WechatMenu struct {
	Key     string         `json:"key"`
	Result  datatypes.JSON `json:"result"`
	AddTime time.Time      `json:"addTIme" gorm:"autoCreateTime"`
}

func (WechatMenu) TableName() string {
	return "wechat_menu"
}

// get all
func GetWechatMenu(maps interface{}) WechatMenu {
	var (
		data WechatMenu
	)

	global.Db.Where(maps).First(&data)

	return data
}

func AddWechatMenu(m *WechatMenu) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatMenu(m *WechatMenu) error {
	var err error
	err = global.Db.Model(&WechatMenu{}).Where("key", constant.WeChatMenu).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}
