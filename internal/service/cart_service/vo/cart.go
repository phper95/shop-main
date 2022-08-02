package vo

import (
	"shop/internal/service/product_service/vo"
)

type Cart struct {
	Id                int64      `json:"id"`
	Uid               int64      `json:"uid"`
	Type              string     `json:"type"`
	ProductId         int64      `json:"product_id"`
	ProductAttrUnique string     `json:"product_attr_unique"`
	CartNum           int        `json:"cart_num"`
	CombinationId     int64      `json:"combination_id"`
	SeckillId         int64      `json:"seckill_id"`
	BargainId         int64      `json:"bargain_id"`
	CostPrice         float64    `json:"cost_price"`
	ProductInfo       vo.Product `json:"product_info"`
	TruePrice         float64    `json:"true_price"`
	TrueStock         int        `json:"true_stock"`
	VipTruePrice      float64    `json:"vip_true_price"`
	Unique            string     `json:"unique"`
	IsReply           int        `json:"is_reply"`
}
