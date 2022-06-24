/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

import "gorm.io/gorm"

type shopUserBill struct {
	Uid      int64   `json:"uid"`
	LinkId   string  `json:"link_id"`
	Pm       int8    `json:"pm"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Type     string  `json:"type"`
	Number   float64 `json:"number"`
	Balance  float64 `json:"balance"`
	Mark     string  `json:"mark"`
	Status   int8    `json:"status"`
	BaseModel
}

func (shopUserBill) TableName() string {
	return "shop_user_bill"
}

//增加支出流水
func Expend(tx *gorm.DB, uid int64, title, category, typestr, mark, linkId string, number, balance float64) error {
	data := &shopUserBill{
		Uid:      uid,
		Title:    title,
		Category: category,
		Type:     typestr,
		Number:   number,
		Balance:  balance,
		Mark:     mark,
		Pm:       0,
		LinkId:   linkId,
	}
	return tx.Model(&shopUserBill{}).Create(data).Error
}

//增加收入流水
func Income(tx *gorm.DB, uid int64, title, category, typestr, mark, linkId string, number, balance float64) error {
	data := &shopUserBill{
		Uid:      uid,
		Title:    title,
		Category: category,
		Type:     typestr,
		Number:   number,
		Balance:  balance,
		Mark:     mark,
		Pm:       1,
		LinkId:   linkId,
	}
	return tx.Model(&shopUserBill{}).Create(data).Error
}

func AddUserBill(m *shopUserBill) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserBill(m *shopUserBill) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserBill(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&shopUserBill{}).Error
	if err != nil {
		return err
	}

	return err
}
