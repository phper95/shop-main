package dto

type UserMoney struct {
	Id    int64   `json:"id"`
	Money float64 `json:"money"`
	Ptype int     `json:"ptype"`
}
