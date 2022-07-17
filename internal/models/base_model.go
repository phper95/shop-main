package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"update_time" gorm:"autoUpdateTime"`
	CreateTime time.Time             `json:"create_time" gorm:"autoCreateTime"`
	IsDel      soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
}
