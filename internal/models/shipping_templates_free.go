package models

import "shop/pkg/global"

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
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplatesFree(m *ShippingTemplatesFree) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatesFree(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&ShippingTemplatesFree{}).Error
	if err != nil {
		return err
	}

	return err
}
