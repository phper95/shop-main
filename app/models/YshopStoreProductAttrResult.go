/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package models

import (
	"encoding/json"
	dto2 "shop/app/service/product_service/dto"
	"shop/pkg/logging"
	"time"
)

type shopStoreProductAttrResult struct {
	ID         int64     `json:"id"`
	ProductId  int64     `json:"productId" valid:"Required;"`
	Result     string    `json:"sliderImage" valid:"Required;"`
	ChangeTime time.Time `json:"change_time" gorm:"autoCreateTime"`
}

func (shopStoreProductAttrResult) TableName() string {
	return "shop_store_product_attr_result"
}

func GetProductAttrResult(productId int64) map[string]interface{} {
	var (
		result shopStoreProductAttrResult
		data   map[string]interface{}
	)
	db.Where("product_id = ?", productId).First(&result)

	e := json.Unmarshal([]byte(result.Result), &data)
	if e != nil {
		logging.Error(e)
	}

	return data
}

func AddProductAttrResult(items []dto2.FormatDetail, attrs []dto2.ProductFormat, productId int64) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var count int64
	mapData := map[string]interface{}{
		"attr":  items,
		"value": attrs,
	}
	b, _ := json.Marshal(mapData)
	db.Model(&shopStoreProductAttrResult{}).Where("product_id = ?", productId).Count(&count)
	if count > 0 {
		err = DelByProductAttrResult(productId)
		if err != nil {
			return err
		}
	}
	var result = shopStoreProductAttrResult{
		ProductId: productId,
		Result:    string(b),
	}

	err = tx.Create(&result).Error
	if err != nil {
		return err
	}
	return err
}

func DelByProductAttrResult(productId int64) (err error) {
	err = db.Where("product_id = ?", productId).Delete(shopStoreProductAttrResult{}).Error
	return err
}
