package models

import (
	"shop/pkg/global"
	"shop/pkg/logging"
	"shop/pkg/util"
)

type SysUser struct {
	Avatar      string     `json:"avatar"`
	Email       string     `json:"email"`
	Enabled     int8       `json:"enabled"`
	Password    string     `json:"password"`
	Username    string     `json:"username" valid:"Required;"`
	DeptId      int32      `json:"deptId"`
	Phone       string     `json:"phone"`
	JobId       int32      `json:"jobId"`
	NickName    string     `json:"nickName"`
	Sex         string     `json:"sex"`
	Roles       []*SysRole `json:"roles" gorm:"many2many:sys_users_roles;association_autoupdate:false;association_autocreate:false"`
	Jobs        *SysJob    `json:"job" gorm:"foreignKey:JobId;association_autoupdate:false;association_autocreate:false"`
	Depts       *SysDept   `json:"dept" gorm:"foreignKey:DeptId;association_autoupdate:false;association_autocreate:false"`
	Permissions []string   `gorm:"-"`
	RoleIds     []int64    `json:"roleIds" gorm:"-"`
	BaseModel
}

type RoleId struct {
	Id int64 `json:"id"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func FindByUserId(id int64) ([]string, error) {
	var roles []SysRole
	var roleIds []int64
	global.Db.Raw("SELECT r.* FROM sys_role r, sys_users_roles u WHERE r.id = u.sys_role_id AND u.sys_user_id = ?", id).Scan(&roles)
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	var rolesMenus []SysRolesMenus
	var menuIds []int64
	global.Db.Table("sys_roles_menus").Where("sys_role_id in (?)", roleIds).Find(&rolesMenus)
	for _, roleMenu := range rolesMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	var menus []SysMenu
	global.Db.Table("sys_menu").Where("id in (?)", menuIds).Find(&menus)

	logging.Info(roles)
	var permissions []string

	for _, m := range menus {
		if m.Permission == "" {
			continue
		}
		permissions = append(permissions, m.Permission)
	}

	return permissions, nil
}

//根据用户名返回
func GetUserByUsername(name string) (*SysUser, error) {
	var user SysUser
	err := global.Db.Preload("Roles").Preload("Jobs").Preload("Depts").Where("username = ? and is_del = ? ", name, 0).First(&user).Error
	if err == nil {
		permissions, _ := FindByUserId(user.Id)
		user.Permissions = permissions
		return &user, nil
	}

	return nil, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (SysUser, error) {
	var user SysUser
	var err error

	err = global.Db.Where("id = ?", id).First(&user).Error

	return user, err
}

// get all
func GetAllUser(pageNUm int, pageSize int, maps interface{}) (int64, []SysUser) {
	var (
		total int64
		users []SysUser
	)

	global.Db.Model(&SysUser{}).Where(maps).Count(&total)
	global.Db.Model(&SysUser{}).Where(maps).Offset(pageNUm).Limit(pageSize).Preload("Jobs").Preload("Depts").Preload("Roles").Find(&users)

	return total, users
}

func UpdateCurrentUser(m *SysUser) (err error) {
	err = global.Db.Save(m).Error
	return
}

func AddUser(m *SysUser) error {
	var err error
	m.Password = util.HashAndSalt([]byte("123456"))
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUser(m *SysUser) error {
	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("sys_user_id = ?", m.Id).Delete(SysUsersRoles{}).Error
	if err != nil {
		return err
	}
	err = tx.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUser(ids []int64) error {
	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Where("id in (?)", ids).Delete(&SysUser{}).Error
	if err != nil {
		return err
	}
	err = tx.Unscoped().Where("sys_user_id in (?)", ids).Delete(SysUsersRoles{}).Error
	if err != nil {
		return err
	}

	return err
}
