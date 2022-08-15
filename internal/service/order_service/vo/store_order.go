package vo

import (
	cartVo "shop/internal/service/cart_service/vo"
	"shop/internal/service/order_service/dto"
	"time"
)

type StoreOrder struct {
	Id                     int64         `json:"id"`
	OrderId                string        `json:"order_id"`
	DisplayOrderId         string        `json:"display_order_id"`
	ExtendOrderId          string        `json:"extend_order_id"`
	Uid                    int64         `json:"uid"`
	RealName               string        `json:"real_name"`
	UserPhone              string        `json:"user_phone"`
	UserAddress            string        `json:"user_address"`
	CartId                 string        `json:"cart_id"`
	CartInfo               []cartVo.Cart `json:"cart_info" copier:"-"`
	StatusDto              dto.Status    `json:"_status"`
	FreightPrice           float64       `json:"freight_price"`
	TotalNum               int           `json:"total_num"`
	TotalPrice             float64       `json:"total_price"`
	TotalPostage           float64       `json:"total_postage"`
	PayPrice               float64       `json:"pay_price"`
	PayPostage             float64       `json:"pay_postage"`
	DeductionPrice         float64       `json:"deduction_price"`
	CouponId               int64         `json:"coupon_id"`
	CouponPrice            float64       `json:"coupon_price"`
	Paid                   int8          `json:"paid"`
	PayTime                time.Time     `json:"pay_time"`
	PayType                string        `json:"pay_type"`
	Status                 int8          `json:"status"`
	RefundStatus           int8          `json:"refund_status"`
	RefundReasonWapImg     string        `json:"refund_reason_map_img"`
	RefundReasonWapWxplain string        `json:"refund_reason_map_explain"`
	RefundReasonTime       time.Time     `json:"refund_reason_time"`
	RefundReasonWap        string        `json:"refund_reason_map"`
	RefundReason           string        `json:"refund_reason"`
	RefundPrice            float64       `json:"refund_price"`
	DeliverySn             string        `json:"delivery_sn"`
	DeliveryName           string        `json:"delivery_name"`
	DeliveryType           string        `json:"delivery_type"`
	DeliveryId             string        `json:"delivery_id"`
	GainIntegral           int           `json:"gain_integral"`
	UseIntegral            int           `json:"use_integral"`
	PayIntegral            int           `json:"pay_integral"`
	BackIntegral           int           `json:"back_integral"`
	Mark                   string        `json:"mark"`
	Remark                 string        `json:"remark"`
	Cost                   float64       `json:"cost"`
	ShippingType           int           `json:"shipping_type"`
	PinkId                 int64         `json:"pink_id"`
	CreateTime             time.Time     `json:"create_time"`
}
