/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

type shopStoreCart struct {
	Uid               int64  `json:"uid"`
	Type              string `json:"type"`
	ProductId         int64  `json:"product_id"`
	ProductAttrUnique string `json:"product_attr_unique"`
	CartNum           int    `json:"cart_num"`
	IsPay             int8   `json:"is_pay"`
	IsNew             int8   `json:"is_new"`
	CombinationId     int64  `json:"combination_id"`
	SeckillId         int64  `json:"seckill_id"`
	BargainId         int64  `json:"bargain_id"`
	BaseModel
}

func (shopStoreCart) TableName() string {
	return "shop_store_cart"
}

// get all
func GetAllStoreCart(pageNUm int, pageSize int, maps interface{}) (int64, []shopStoreCart) {
	var (
		total int64
		data  []shopStoreCart
	)
	db.Model(&shopStoreCart{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddStoreCart(m *shopStoreCart) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreCart(m *shopStoreCart) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreCart(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&shopStoreCart{}).Error
	if err != nil {
		return err
	}

	return err
}
