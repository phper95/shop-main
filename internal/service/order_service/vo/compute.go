package vo

type Compute struct {
	CouponPrice    float64 `json:"coupon_price"`
	DeductionPrice float64 `json:"deduction_price"`
	PayPostage     float64 `json:"pay_postage"`
	PayPrice       float64 `json:"pay_price"`
	TotalPrice     float64 `json:"total_price"`
	UseIntegral    int     `json:"use_integral"`
	PayIntegral    int     `json:"pay_integral"`
}
