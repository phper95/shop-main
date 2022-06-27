package vo

type Compute struct {
	CouponPrice    float64 `json:"couponPrice"`
	DeductionPrice float64 `json:"deductionPrice"`
	PayPostage     float64 `json:"payPostage"`
	PayPrice       float64 `json:"payPrice"`
	TotalPrice     float64 `json:"totalPrice"`
	UseIntegral    int     `json:"useIntegral"`
	PayIntegral    int     `json:"payIntegral"`
}
