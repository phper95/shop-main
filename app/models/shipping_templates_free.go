/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

type ShippingTemplatesFree struct {
	Id         int64   `gorm:"primary_key" json:"id"`
	ProvinceId int     `json:"province_id"`
	TempId     int64   `json:"temp_id"`
	CityId     int     `json:"city_id"`
	Number     float64 `json:"number"`
	Price      float64 `json:"price"`
	Type       int8    `json:"type"`
	Uniqid     string  `json:"uniqid"`
}

func (ShippingTemplatesFree) TableName() string {
	return "shipping_templates_free"
}

func AddShippingTemplatesFree(m *ShippingTemplatesFree) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplatesFree(m *ShippingTemplatesFree) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatesFree(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&ShippingTemplatesFree{}).Error
	if err != nil {
		return err
	}

	return err
}
