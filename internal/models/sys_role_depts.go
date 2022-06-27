package models

type SysRolesDepts struct {
	Id     int64
	RoleId *SysRole `gorm:"column:sys_role_id;association_autoupdate:false;association_autocreate:false"`
	DeptId *SysDept `gorm:"column:sys_dept_id;association_autoupdate:false;association_autocreate:false"`
}
