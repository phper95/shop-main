package models

import "shop/pkg/global"

type SysRole struct {
	Name       string     `json:"name" valid:"Required;"`
	Remark     string     `json:"remark"`
	DataScope  string     `json:"dataScope"`
	Level      int32      `json:"level"`
	Permission string     `json:"permission"`
	Users      []*SysUser `gorm:"many2many:sys_users_roles;association_autoupdate:false;association_autocreate:false"`
	Menus      []*SysMenu `json:"menus" gorm:"many2many:sys_roles_menus;association_autoupdate:false;association_autocreate:false"`
	Depts      []*SysDept `gorm:"many2many:sys_roles_depts;association_autoupdate:false;association_autocreate:false"`
	BaseModel
}

func (SysRole) TableName() string {
	return "sys_role"
}

func GetOneRole(id int64) SysRole {
	var role SysRole
	global.Db.Where("id = ?", id).First(&role)
	return role
}

// get all
func GetAllRole(pageNUm int, pageSize int, maps interface{}) (int64, []SysRole) {
	var (
		total int64
		lists []SysRole
	)

	global.Db.Model(&SysRole{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("Menus").Find(&lists)

	return total, lists
}

func AddRole(m *SysRole) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByRole(m *SysRole) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByRole(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&SysRole{}).Error
	if err != nil {
		return err
	}

	return err
}
