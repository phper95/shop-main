package params

import (
	"github.com/astaxie/beego/validation"
	"shop/pkg/global"
)

type CartParam struct {
	ProductId     int64  `json:"product_id"`
	UniqueId      string `json:"unique_id"`
	CartNum       int    `json:"cart_num"`
	IsNew         int8   `json:"is_new"`
	CombinationId int64  `json:"combination_id"`
	SeckillId     int64  `json:"seckill_id"`
	BargainId     int64  `json:"bargain_id"`
}

func (p *CartParam) Valid(v *validation.Validation) {
	global.LOG.Info(p.CartNum)
	if vv := v.Range(p.CartNum, 1, 999, "购物车数量"); !vv.Ok {
		vv.Message("数量只能1-999之间")
		return
	}
	if vv := v.Required(p.ProductId, "shop-warning"); !vv.Ok {
		vv.Message("参数有误")
		return
	}
}
