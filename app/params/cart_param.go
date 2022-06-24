/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package params

import (
	"github.com/astaxie/beego/validation"
	"shop/pkg/global"
)

type CartParam struct {
	ProductId     int64  `json:"productId"`
	UniqueId      string `json:"uniqueId"`
	CartNum       int    `json:"cartNum"`
	IsNew         int8   `json:"isNew"`
	CombinationId int64  `json:"combinationId"`
	SeckillId     int64  `json:"seckillId"`
	BargainId     int64  `json:"bargainId"`
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
