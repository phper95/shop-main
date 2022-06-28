package models

import "shop/pkg/global"

type ShippingTemplates struct {
	Name        string `json:"name"`
	Type        int8   `json:"type"`
	RegionInfo  string `json:"region_info"`
	Appoint     int8   `json:"appoint"`
	AppointInfo string `json:"appoint_info"`
	Sort        int8   `json:"sort"`
	BaseModel
}

func (ShippingTemplates) TableName() string {
	return "shipping_templates"
}

func AddShippingTemplates(m *ShippingTemplates) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplates(m *ShippingTemplates) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatess(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&ShippingTemplates{}).Error
	if err != nil {
		return err
	}

	return err
}
