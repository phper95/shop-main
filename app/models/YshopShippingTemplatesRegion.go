/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

type shopShippingTemplatesRegion struct {
	Id            int64   `gorm:"primary_key" json:"id"`
	ProvinceId    int     `json:"province_id"`
	TempId        int64   `json:"temp_id"`
	CityId        int     `json:"city_id"`
	First         float64 `json:"first"`
	FirstPrice    float64 `json:"first_price"`
	Continues     float64 `json:"continues"`
	ContinuePrice float64 `json:"continue_price"`
	Type          int8    `json:"type"`
	Uniqid        string  `json:"uniqid"`
}

func (shopShippingTemplatesRegion) TableName() string {
	return "shop_shipping_templates_region"
}

func AddShippingTemplatesRegion(m *shopShippingTemplatesRegion) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplatesRegion(m *shopShippingTemplatesRegion) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatesRegion(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&shopShippingTemplatesRegion{}).Error
	if err != nil {
		return err
	}

	return err
}
