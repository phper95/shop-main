package params

import (
	"github.com/astaxie/beego/validation"
)

type CartIdsParam struct {
	Ids []int64 `json:"ids"`
}

func (p *CartIdsParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.Ids, "shop-warning"); !vv.Ok {
		vv.Message("参数有误")
		return
	}
}
