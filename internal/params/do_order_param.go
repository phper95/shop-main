package params

import (
	"github.com/astaxie/beego/validation"
)

type DoOrderParam struct {
	Uni string `json:"uni"`
}

func (p *DoOrderParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.Uni, "shop-warning"); !vv.Ok {
		vv.Message("参数有误")
		return
	}
}
