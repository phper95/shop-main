package models

import (
	"encoding/json"
	dto2 "shop/internal/service/product_service/dto"
	"shop/pkg/global"
)

//

type StoreProductRule struct {
	RuleName  string `json:"ruleName" valid:"Required;"`
	RuleValue string `json:"ruleValue" valid:"Required;"`
	BaseModel
}

func (StoreProductRule) TableName() string {
	return "store_product_rule"
}

// get all
func GetAllProductRule(pageNUm int, pageSize int, maps interface{}) (int64, []dto2.ProductRule) {
	var (
		total   int64
		data    []StoreProductRule
		retData []dto2.ProductRule
	)
	global.Db.Model(&StoreProductRule{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	for _, rule := range data {
		var value []interface{}
		json.Unmarshal([]byte(rule.RuleValue), &value)
		v := dto2.ProductRule{
			Id:         rule.Id,
			RuleName:   rule.RuleName,
			RuleValue:  value,
			CreateTime: rule.CreateTime,
		}

		retData = append(retData, v)
	}

	return total, retData
}

func AddProductRule(m *StoreProductRule) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByProductRule(id int64, m *StoreProductRule) error {
	var err error
	err = global.Db.Model(&StoreProductRule{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByProductRulee(ids []int64) error {
	var err error
	err = global.Db.Model(&StoreProductRule{}).Where("id in (?)", ids).Update("is_del", 1).Error
	if err != nil {
		return err
	}

	return err
}
