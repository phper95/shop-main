package vo

import (
	"shop/internal/models"
)

type ProductDetail struct {
	ProductAttr  []ProductAttr                           `json:"product_attr"`
	ProductValue map[string]models.StoreProductAttrValue `json:"product_value"`
	Reply        models.StoreProductReply                `json:"reply"`
	ReplyChance  string                                  `json:"reply_chance"`
	ReplyCount   string                                  `json:"reply_count"`
	StoreInfo    Product                                 `json:"store_info"`
	Uid          int64                                   `json:"uid"`
	TempName     string                                  `json:"temp_name"`
}
