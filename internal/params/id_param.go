package params

import (
	"github.com/astaxie/beego/validation"
)

type IdParam struct {
	Id int64 `json:"id"`
}

func (p *IdParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.Id, "shop-warning"); !vv.Ok {
		vv.Message("参数有误")
		return
	}
}
