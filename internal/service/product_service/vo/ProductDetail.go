package vo

import (
	"shop/internal/models"
)

type ProductDetail struct {
	ProductAttr  []ProductAttr                           `json:"productAttr"`
	ProductValue map[string]models.StoreProductAttrValue `json:"productValue"`
	Reply        models.StoreProductReply                `json:"reply"`
	ReplyChance  string                                  `json:"replyChance"`
	ReplyCount   string                                  `json:"replyCount"`
	StoreInfo    Product                                 `json:"storeInfo"`
	Uid          int64                                   `json:"uid"`
	TempName     string                                  `json:"tempName"`
}
