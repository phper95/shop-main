/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

import (
	"gorm.io/datatypes"
	"shop/pkg/constant"
	"time"
)

type WechatMenu struct {
	Key     string         `json:"key"`
	Result  datatypes.JSON `json:"result"`
	AddTime time.Time      `json:"addTIme" gorm:"autoCreateTime"`
}

func (WechatMenu) TableName() string {
	return "shop_wechat_menu"
}

// get all
func GetWechatMenu(maps interface{}) WechatMenu {
	var (
		data WechatMenu
	)

	db.Where(maps).First(&data)

	return data
}

func AddWechatMenu(m *WechatMenu) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatMenu(m *WechatMenu) error {
	var err error
	err = db.Model(&WechatMenu{}).Where("key", constant.Shop_WEICHAT_MENU).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}
