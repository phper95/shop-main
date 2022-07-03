package models

import "shop/pkg/global"

type UserAddress struct {
	Uid       int64  `json:"uid"`
	RealName  string `json:"realName"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	CityId    int    `json:"cityId"`
	District  string `json:"district"`
	Detail    string `json:"detail"`
	PostCode  string `json:"postCode"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	IsDefault int8   `json:"isDefault" gorm:"force"`
	BaseModel
}

func (UserAddress) TableName() string {
	return "user_address"
}

// get all
func GetAllUserAddress(pageNUm int, pageSize int, maps interface{}) (int64, []UserAddress) {
	var (
		total int64
		data  []UserAddress
	)

	global.Db.Model(&UserAddress{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddUserAddress(m *UserAddress) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserAddress(m *UserAddress) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserAddress(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&UserAddress{}).Error
	if err != nil {
		return err
	}

	return err
}
