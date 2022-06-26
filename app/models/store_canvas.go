/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

type StoreCanvas struct {
	Terminal int    `json:"terminal"`
	Json     string `json:"json"`
	Ttype    int    `json:"type" gorm:"column:type"`
	Name     string `json:"name"`
	BaseModel
}

func (StoreCanvas) TableName() string {
	return "store_canvas"
}

func AddCanvas(m *StoreCanvas) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByCanvas(m *StoreCanvas) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByCanvas(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&StoreCanvas{}).Error
	if err != nil {
		return err
	}

	return err
}
