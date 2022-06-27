package vo

type ProductReply struct {
	Uid          int64    `json:"uid"`
	ProductId    int64    `json:"product_id"`
	Oid          int64    `json:"oid"`
	Unique       string   `json:"unique"`
	ReplyType    string   `json:"reply_type"`
	ProductScore int      `json:"product_score"`
	ServiceScore int      `json:"service_score"`
	Comment      string   `json:"comment"`
	PicturesArr  []string `json:"picturesArr"`
}
