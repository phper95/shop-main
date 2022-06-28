package models

import "shop/pkg/global"

type SysJob struct {
	Name    string   `json:"name" valid:"Required;"`
	Enabled int8     `json:"enabled"`
	Sort    int      `json:"sort"`
	DeptId  int64    `json:"deptId"`
	Dept    *SysDept `json:"dept" gorm:"foreignKey:DeptId;association_autoupdate:false;association_autocreate:false"`
	BaseModel
}

func (SysJob) TableName() string {
	return "sys_job"
}

// get all
func GetAllJob(pageNUm int, pageSize int, maps interface{}) (int64, []SysJob) {
	var (
		total int64
		lists []SysJob
	)

	global.Db.Model(&SysJob{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("Dept").Find(&lists)

	return total, lists
}

func AddJob(m *SysJob) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByJob(m *SysJob) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByJob(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&SysJob{}).Error
	if err != nil {
		return err
	}

	return err
}
