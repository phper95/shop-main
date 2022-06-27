package params

type ReplyParam struct {
	Comment      string `json:"comment"`
	Pics         string `json:"pics"`
	ProductScore int    `json:"productScore"`
	ServiceScore int    `json:"serviceScore"`
	Unique       string `json:"unique"`
	ProductId    int64  `json:"productId"`
}
