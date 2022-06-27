package material_group_service

import (
	"shop/internal/models"
)

type MaterialGroup struct {
	Id   int64
	Name string

	Enabled int
	GroupId int64

	PageNum  int
	PageSize int

	M *models.SysMaterialGroup

	Ids []int64
}

func (d *MaterialGroup) GetAll() []models.SysMaterialGroup {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	list := models.GetAllGroup(maps)
	return list
}

func (d *MaterialGroup) Insert() error {
	return models.AddMaterialGroup(d.M)
}

func (d *MaterialGroup) Save() error {
	return models.UpdateByMaterialGroup(d.M)
}

func (d *MaterialGroup) Del() error {
	return models.DelByMaterialGroup(d.Ids)
}
