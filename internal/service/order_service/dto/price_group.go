package dto

type PriceGroup struct {
	CostPrice        float64 `json:"cost_price"`
	StoreFreePostage float64 `json:"store_free_postage"`
	StorePostage     float64 `json:"store_postage"`
	TotalPrice       float64 `json:"total_price"`
	VipPrice         float64 `json:"vip_price"`
	PayIntegral      int     `json:"pay_integral"`
}
