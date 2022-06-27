package vo

import "shop/internal/models"

type LoginVo struct {
	Token string          `json:"token"`
	User  *models.SysUser `json:"user"`
}
