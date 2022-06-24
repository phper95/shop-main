/**
¬* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
*/
package models

type shopStoreProductRelation struct {
	Uid       int64             `json:"uid"`
	ProductId int64             `json:"productId"`
	Type      string            `json:"type"`
	Category  string            `json:"category"`
	Product   *shopStoreProduct `json:"product" gorm:"foreignKey:ProductId;"`
	BaseModel
}

func (shopStoreProductRelation) TableName() string {
	return "shop_store_product_relation"
}

// get all
func GetAllProductRelation(pageNUm int, pageSize int, maps interface{}) (int64, []shopStoreProductRelation) {
	var (
		total int64
		data  []shopStoreProductRelation
	)
	db.Model(&shopStoreProductRelation{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("Product").Order("id desc").Find(&data)

	return total, data
}

func AddStoreProductRelation(m *shopStoreProductRelation) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreProductRelation(m *shopStoreProductRelation) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreProductRelations(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&shopStoreProductRelation{}).Error
	if err != nil {
		return err
	}

	return err
}
