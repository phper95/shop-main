package vo

import (
	"shop/internal/models"
	"shop/internal/service/cart_service/vo"
	dto2 "shop/internal/service/order_service/dto"
	vo2 "shop/internal/service/wechat_user_service/vo"
)

type ConfirmOrder struct {
	AddressInfo       models.UserAddress `json:"address_info"`
	BargainId         int64              `json:"bargain_id"`
	CartInfo          []vo.Cart          `json:"cart_info"`
	CombinationId     int64              `json:"combination_id"`
	Deduction         bool               `json:"deduction"`
	EnableIntegral    bool               `json:"enable_integral"`
	SeckillId         int64              `json:"seckill_id"`
	EnableIntegralNum int                `json:"enable_integral_num"`
	IntegralRadio     int                `json:"integral_radio"`
	OrderKey          string             `json:"order_key"`
	StoreSelfMention  int                `json:"store_self_mention"`
	//UsableCoupon string `json:"usableCoupon"`
	//SystemStore string `json:"systemStore"`
	UserInfo   vo2.User        `json:"user_info"`
	PriceGroup dto2.PriceGroup `json:"price_group"`
}
