package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime time.Time             `json:"createTime" gorm:"autoCreateTime"`
	IsDel      soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`
}
