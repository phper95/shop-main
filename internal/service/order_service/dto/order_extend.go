package dto

type OrderExtend struct {
	Key      string                 `json:"key"`
	OrderId  string                 `json:"order_id"`
	JsConfig map[string]interface{} `json:"js_config"`
}
