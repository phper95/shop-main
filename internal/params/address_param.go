package params

import "github.com/astaxie/beego/validation"

type AddressParan struct {
	Id        int64              `json:"id"`
	RealName  string             `json:"real_name"`
	Phone     string             `json:"phone"`
	Detail    string             `json:"detail"`
	PostCode  string             `json:"post_code"`
	IsDefault bool               `json:"is_default"`
	Address   AddressDetailParan `json:"address"`
}

func (p *AddressParan) Valid(v *validation.Validation) {
	if vv := v.MaxSize(p.RealName, 30, "姓名"); !vv.Ok {
		vv.Message("姓名不能超过30")
		return
	}
	if vv := v.MaxSize(p.Detail, 60, "姓名"); !vv.Ok {
		vv.Message("姓名不能超过60")
		return
	}
}
