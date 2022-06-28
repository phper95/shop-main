package models

import "shop/pkg/global"

type Express struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	IsShow int8   `json:"is_show"`
	BaseModel
}

func (Express) TableName() string {
	return "express"
}

// get all
func GetAllExpress(pageNUm int, pageSize int, maps interface{}) (int64, []Express) {
	var (
		total int64
		lists []Express
	)

	global.Db.Model(&Express{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists
}

func AddExpress(m *Express) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByExpress(m *Express) error {
	var err error
	err = global.Db.Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByExpress(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&Express{}).Error
	if err != nil {
		return err
	}

	return err
}
