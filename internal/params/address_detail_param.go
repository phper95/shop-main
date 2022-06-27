package params

type AddressDetailParan struct {
	Province string `json:"province"`
	City     string `json:"city"`
	CityId   int    `json:"cityId"`
	District string `json:"district"`
}
