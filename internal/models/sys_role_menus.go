package models

import (
	dto2 "shop/internal/service/menu_service/dto"
	"shop/pkg/constant"
	"shop/pkg/global"
	"shop/pkg/runtime"
)

type SysRolesMenus struct {
	ID     int64 `gorm:"primaryKey;autoIncrement"`
	MenuId int64 `gorm:"column:sys_menu_id;"`
	RoleId int64 `gorm:"column:sys_role_id;"`
}

func BatchRoleMenuAdd(menu dto2.RoleMenu) error {

	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Where("sys_role_id = ?", menu.Id).Delete(SysRolesMenus{}).Error
	if err != nil {
		return err
	}

	var roleMenus []SysRolesMenus
	var roles = GetOneRole(menu.Id)

	cb := runtime.Runtime.GetCasbinKey(constant.CASBIN)
	cb.RemoveFilteredPolicy(0, roles.Permission)
	for _, val := range menu.Menus {
		var roleMenu = SysRolesMenus{RoleId: menu.Id, MenuId: val.Id}

		var menus = GetOneMenuById(val.Id)
		roleMenus = append(roleMenus, roleMenu)
		if roles.Permission != "" && menus.Router != "" && menus.RouterMethod != "" {
			cb.AddNamedPolicy("p", roles.Permission, menus.Router, menus.RouterMethod)
		}

	}

	err = tx.Create(&roleMenus).Error
	if err != nil {
		return err
	}
	//logging.Info(roleMenus)
	cb.SavePolicy()

	return err
}
