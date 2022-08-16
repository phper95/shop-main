package models

import (
	"shop/pkg/global"
	"time"
)

type StoreOrder struct {
	OrderId                string      `json:"order_id"`
	ExtendOrderId          string      `json:"extend_order_id"`
	Uid                    int64       `json:"uid"`
	RealName               string      `json:"real_name"`
	UserPhone              string      `json:"user_phone"`
	UserAddress            string      `json:"user_address"`
	CartId                 string      `json:"cart_id"`
	FreightPrice           float64     `json:"freight_price"`
	TotalNum               int         `json:"total_num"`
	TotalPrice             float64     `json:"total_price"`
	TotalPostage           float64     `json:"total_postage"`
	PayPrice               float64     `json:"pay_price"`
	PayPostage             float64     `json:"pay_postage"`
	DeductionPrice         float64     `json:"deduction_price"`
	CouponId               int64       `json:"coupon_id"`
	CouponPrice            float64     `json:"coupon_price"`
	Paid                   int         `json:"paid"`
	PayTime                time.Time   `json:"pay_time"`
	PayType                string      `json:"pay_type"`
	Status                 int         `json:"status"`
	RefundStatus           int         `json:"refund_status"`
	RefundReasonWapImg     string      `json:"refund_reason_wap_img"`
	RefundReasonWapExplain string      `json:"refund_reason_wap_explain"`
	RefundReasonTime       time.Time   `json:"refund_reason_time"`
	RefundReasonWap        string      `json:"refund_reason_wap"`
	RefundReason           string      `json:"refund_reason"`
	RefundPrice            float64     `json:"refund_price"`
	DeliverySn             string      `json:"delivery_sn"`
	DeliveryName           string      `json:"delivery_name"`
	DeliveryType           string      `json:"delivery_type"`
	DeliveryId             string      `json:"delivery_id"`
	GainIntegral           int         `json:"gain_integral"`
	UseIntegral            int         `json:"use_integral"`
	PayIntegral            int         `json:"pay_integral"`
	BackIntegral           int         `json:"back_integral"`
	Mark                   string      `json:"mark"`
	Unique                 string      `json:"unique"`
	Remark                 string      `json:"remark"`
	CombinationId          int64       `json:"combination_id"`
	PinkId                 int64       `json:"pink_id"`
	Cost                   float64     `json:"cost"`
	SeckillId              int64       `json:"seckill_id"`
	BargainId              int64       `json:"bargain_id"`
	VerifyCode             string      `json:"verify_code"`
	StoreId                int64       `json:"store_id"`
	ShippingType           int         `json:"shipping_type"`
	UserDto                *ShopUser   `json:"user_dto" gorm:"foreignKey:Uid;"`
	CartInfo               interface{} `json:"cart_info" gorm:"-" copier:"-"`
	OrderStatus            int         `json:"_status" gorm:"-"`
	OrderStatusName        string      `json:"status_name" gorm:"-"`
	PayTypeName            string      `json:"pay_type_name" gorm:"-"`
	BaseModel
}

//定义订单消息结构
type OrderMsg struct {
	Operation string `json:"operation"`
	*StoreOrder
}

func (StoreOrder) TableName() string {
	return "store_order"
}

func GetAllOrder(pageNUm int, pageSize int, maps interface{}) (int64, []StoreOrder) {
	var (
		total int64
		data  []StoreOrder
	)

	global.Db.Model(&StoreOrder{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func GetAdminAllOrder(pageNUm int, pageSize int, maps interface{}) (int64, []StoreOrder) {
	var (
		total int64
		data  []StoreOrder
	)

	global.Db.Model(&StoreOrder{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("UserDto").Order("id desc").Find(&data)

	return total, data
}
func GetOrderUseCursor(userID, nextID int64, pageSize int) []StoreOrder {
	var (
		data []StoreOrder
	)
	//SELECT * FROM `store_order` WHERE (uid = 4 AND id > 0) AND `store_order`.`is_del` = 0 ORDER BY id asc LIMIT 10
	global.Db.Where("uid = ? AND id > ?", userID, nextID).Limit(pageSize).Preload("UserDto").Order("id asc").Find(&data)

	return data
}

func AddStoreOrder(m *StoreOrder) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreOrder(m *StoreOrder) error {
	var err error
	err = global.Db.Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreOrder(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&StoreOrder{}).Error
	if err != nil {
		return err
	}

	return err
}
