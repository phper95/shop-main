package models

import (
	dto2 "shop/internal/service/product_service/dto"
	"shop/pkg/global"
	"strings"
)

type StoreProductAttr struct {
	Id         int64  `json:"id"`
	ProductId  int64  `json:"product_id" valid:"Required;"`
	AttrName   string `json:"attr_name" valid:"Required;"`
	AttrValues string `json:"attr_values" valid:"Required;"`
}

func (StoreProductAttr) TableName() string {
	return "store_product_attr"
}

func AddProductAttr(items []dto2.FormatDetail, productId int64) error {
	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var attrGroup []StoreProductAttr
	for _, val := range items {
		detailStr := strings.Join(val.Detail, ",")
		var storeProductAttr = StoreProductAttr{
			ProductId:  productId,
			AttrName:   val.Value,
			AttrValues: detailStr,
		}
		attrGroup = append(attrGroup, storeProductAttr)
	}

	err = tx.Create(&attrGroup).Error
	if err != nil {
		return err
	}

	return err
}

func DelByProductttr(productId int64) (err error) {
	err = global.Db.Where("product_id = ?", productId).Delete(StoreProductAttr{}).Error
	return err
}
