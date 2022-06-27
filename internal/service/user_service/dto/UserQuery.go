package dto

import "shop/internal/models/dto"

type UserQuery struct {
	dto.BasePage
	Sort    string
	Blurry  string
	Enabled bool
}
