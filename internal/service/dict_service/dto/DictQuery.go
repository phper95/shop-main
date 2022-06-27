package dto

import "shop/internal/models/dto"

type DictQuery struct {
	dto.BasePage
	Blurry string
}
