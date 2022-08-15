package dto

import (
	"shop/internal/service/cart_service/vo"
)

type Cache struct {
	CartInfo   []vo.Cart  `json:"cart_info"`
	PriceGroup PriceGroup `json:"price_group"`
	Other      Other      `json:"other"`
}
