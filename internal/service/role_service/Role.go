package role_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
	dto2 "shop/internal/service/menu_service/dto"
)

type Role struct {
	Id   int64
	Name string

	PageNum  int
	PageSize int

	M *models.SysRole

	Ids []int64

	Dto dto2.RoleMenu
}

func (d *Role) GetOneRole() models.SysRole {
	return models.GetOneRole(d.Id)
}

func (d *Role) BatchRoleMenuAdd() error {
	return models.BatchRoleMenuAdd(d.Dto)
}

func (d *Role) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllRole(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Role) Insert() error {
	return models.AddRole(d.M)
}

func (d *Role) Save() error {
	return models.UpdateByRole(d.M)
}

func (d *Role) Del() error {
	return models.DelByRole(d.Ids)
}
