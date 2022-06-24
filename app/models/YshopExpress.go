/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

type shopExpress struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	IsShow int8   `json:"is_show"`
	BaseModel
}

func (shopExpress) TableName() string {
	return "shop_express"
}

// get all
func GetAllExpress(pageNUm int, pageSize int, maps interface{}) (int64, []shopExpress) {
	var (
		total int64
		lists []shopExpress
	)

	db.Model(&shopExpress{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists
}

func AddExpress(m *shopExpress) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByExpress(m *shopExpress) error {
	var err error
	err = db.Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByExpress(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&shopExpress{}).Error
	if err != nil {
		return err
	}

	return err
}
