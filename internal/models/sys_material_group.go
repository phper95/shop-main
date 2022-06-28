package models

import "shop/pkg/global"

type SysMaterialGroup struct {
	Name     string `json:"name" valid:"Required;"`
	CreateId int64  `json:"create_id"`
	BaseModel
}

func (SysMaterialGroup) TableName() string {
	return "sys_material_group"
}

//
func GetAllGroup(maps interface{}) []SysMaterialGroup {
	var data []SysMaterialGroup
	global.Db.Where(maps).Find(&data)
	return data
}

func AddMaterialGroup(m *SysMaterialGroup) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByMaterialGroup(m *SysMaterialGroup) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByMaterialGroup(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&SysMaterialGroup{}).Error
	if err != nil {
		return err
	}

	return err
}
