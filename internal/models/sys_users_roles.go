package models

type SysUsersRoles struct {
	Id     int64
	UserId int64 `gorm:"column:sys_user_id;"`
	RoleId int64 `gorm:"column:sys_role_id;"`
}
