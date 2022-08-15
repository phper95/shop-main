package params

import (
	"github.com/astaxie/beego/validation"
)

type ProductReplyParam struct {
	Comment      string `json:"comment"`
	Pics         string `json:"pics"`
	ProductScore int    `json:"product_score"`
	ServiceScore int    `json:"service_score"`
	Unique       string `json:"unique"`
	ProductId    int64  `json:"product_id"`
}

func (p *ProductReplyParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.Comment, "shop-warning"); !vv.Ok {
		vv.Message("请填写评价内容")
		return
	}
}
