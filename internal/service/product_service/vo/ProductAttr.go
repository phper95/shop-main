package vo

import (
	dto2 "shop/internal/service/product_service/dto"
)

type ProductAttr struct {
	Id           int64            `json:"id"`
	ProductId    int64            `json:"product_id"`
	AttrName     string           `json:"attr_name"`
	AttrValues   string           `json:"attr_values"`
	AttrValue    []dto2.AttrValue `json:"attr_value"`
	AttrValueArr []string         `json:"attr_value_arr"`
}
