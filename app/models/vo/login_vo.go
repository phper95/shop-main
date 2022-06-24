package vo

import "shop/app/models"

type LoginVo struct {
	Token string          `json:"token"`
	User  *models.SysUser `json:"user"`
}
