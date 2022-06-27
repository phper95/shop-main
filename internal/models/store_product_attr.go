package models

import (
	dto2 "shop/internal/service/product_service/dto"
	"strings"
)

type StoreProductAttr struct {
	Id         int64  `json:"id"`
	ProductId  int64  `json:"productId" valid:"Required;"`
	AttrName   string `json:"attrName" valid:"Required;"`
	AttrValues string `json:"attrValues" valid:"Required;"`
}

func (StoreProductAttr) TableName() string {
	return "store_product_attr"
}

func AddProductAttr(items []dto2.FormatDetail, productId int64) error {
	var err error
	tx := db.Begin()
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
	err = db.Where("product_id = ?", productId).Delete(StoreProductAttr{}).Error
	return err
}
