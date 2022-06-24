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

type shopWechatMenu struct {
	Key     string         `json:"key"`
	Result  datatypes.JSON `json:"result"`
	AddTime time.Time      `json:"addTIme" gorm:"autoCreateTime"`
}

func (shopWechatMenu) TableName() string {
	return "shop_wechat_menu"
}

// get all
func GetWechatMenu(maps interface{}) shopWechatMenu {
	var (
		data shopWechatMenu
	)

	db.Where(maps).First(&data)

	return data
}

func AddWechatMenu(m *shopWechatMenu) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatMenu(m *shopWechatMenu) error {
	var err error
	err = db.Model(&shopWechatMenu{}).Where("key", constant.shop_WEICHAT_MENU).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}
