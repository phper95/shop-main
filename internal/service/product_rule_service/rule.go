package product_rule_service

import (
	"encoding/json"
	"shop/internal/models"
	"shop/internal/models/vo"
	dto2 "shop/internal/service/product_service/dto"
)

type Rule struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.StoreProductRule

	Ids []int64

	Dto dto2.ProductRule
}

func (d *Rule) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllProductRule(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Rule) AddOrSave() error {
	jsonstr, _ := json.Marshal(d.Dto.RuleValue)
	ruleValue := string(jsonstr)
	if d.Id > 0 {
		model := &models.StoreProductRule{
			RuleName:  d.Dto.RuleName,
			RuleValue: ruleValue,
		}
		return models.UpdateByProductRule(d.Id, model)
	} else {
		model := &models.StoreProductRule{
			RuleName:  d.Dto.RuleName,
			RuleValue: ruleValue,
		}
		return models.AddProductRule(model)
	}

}

func (d *Rule) Insert() error {
	return models.AddProductRule(d.M)
}

func (d *Rule) Save() error {
	return models.UpdateByProductRule(d.Id, d.M)
}

func (d *Rule) Del() error {
	return models.DelByProductRulee(d.Ids)
}
