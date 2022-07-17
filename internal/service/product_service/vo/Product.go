package vo

import "shop/internal/models"

type Product struct {
	Id             int64                        `json:"id"`
	Image          string                       `json:"image"`
	SliderImage    string                       `json:"slider_image"`
	SliderImageArr []string                     `json:"slider_image_arr"`
	StoreName      string                       `json:"store_name"`
	StoreInfo      string                       `json:"store_info"`
	Keyword        string                       `json:"keyword"`
	CateId         int                          `json:"cate_id"`
	Price          float64                      `json:"price"`
	VipPrice       float64                      `json:"vip_price"`
	OtPrice        float64                      `json:"ot_price"`
	Postage        float64                      `json:"postage"`
	UnitName       string                       `json:"unit_name"`
	Sort           int16                        `json:"sort"`
	Sales          int                          `json:"sales"`
	Stock          int                          `json:"stock"`
	Description    string                       `json:"description"`
	IsPostage      int8                         `json:"is_postage"`
	GiveIntegral   int                          `json:"give_integral"`
	Cost           float64                      `json:"cost"`
	IsGood         int8                         `json:"is_good"`
	Ficti          int                          `json:"ficti"`
	Browse         int                          `json:"browse"`
	IsSub          int8                         `json:"is_sub"`
	TempId         int64                        `json:"temp_id"`
	SpecType       int8                         `json:"spec_type"`
	IsIntegral     int8                         `json:"is_integral"`
	Integral       int32                        `json:"integral"`
	UserCollect    bool                         `json:"user_collect"`
	AttrInfo       models.StoreProductAttrValue `json:"attr_info"`
}
