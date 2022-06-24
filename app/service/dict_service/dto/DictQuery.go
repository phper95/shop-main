package dto

import "shop/app/models/dto"

type DictQuery struct {
	dto.BasePage
	Blurry string
}