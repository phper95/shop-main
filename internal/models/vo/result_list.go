package vo

type ResultList struct {
	Content       interface{} `json:"content"`
	TotalElements int64       `json:"totalElements"`
	ExtendData    interface{} `json:"extendData"`
}

type CursorResultList struct {
	Content    interface{} `json:"content"`
	NextID     int64       `json:"next_id"`
	ExtendData interface{} `json:"extendData"`
}
