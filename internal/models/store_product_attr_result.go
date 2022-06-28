package models

import (
	"encoding/json"
	dto2 "shop/internal/service/product_service/dto"
	"shop/pkg/global"
	"shop/pkg/logging"
	"time"
)

type StoreProductAttrResult struct {
	ID         int64     `json:"id"`
	ProductId  int64     `json:"productId" valid:"Required;"`
	Result     string    `json:"sliderImage" valid:"Required;"`
	ChangeTime time.Time `json:"change_time" gorm:"autoCreateTime"`
}

func (StoreProductAttrResult) TableName() string {
	return "store_product_attr_result"
}

func GetProductAttrResult(productId int64) map[string]interface{} {
	var (
		result StoreProductAttrResult
		data   map[string]interface{}
	)
	global.Db.Where("product_id = ?", productId).First(&result)

	e := json.Unmarshal([]byte(result.Result), &data)
	if e != nil {
		logging.Error(e)
	}

	return data
}

func AddProductAttrResult(items []dto2.FormatDetail, attrs []dto2.ProductFormat, productId int64) error {
	var err error
	tx := global.Db.Begin()
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
	global.Db.Model(&StoreProductAttrResult{}).Where("product_id = ?", productId).Count(&count)
	if count > 0 {
		err = DelByProductAttrResult(productId)
		if err != nil {
			return err
		}
	}
	var result = StoreProductAttrResult{
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
	err = global.Db.Where("product_id = ?", productId).Delete(StoreProductAttrResult{}).Error
	return err
}
