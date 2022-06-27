package dept_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
)

type Dept struct {
	Id      int64
	Name    string
	Enabled int

	M *models.SysDept

	Ids []int64
}

func (d *Dept) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Enabled >= 0 {
		maps["enabled"] = d.Enabled
	}
	if d.Name != "" {
		maps["name"] = d.Name
	}

	list := models.GetAllDepts(maps)
	return vo.ResultList{Content: list, TotalElements: 0}
}

func (d *Dept) Insert() error {
	return models.AddDept(d.M)
}

func (d *Dept) Save() error {
	return models.UpdateByDept(d.M)
}

func (d *Dept) Del() error {
	return models.DelByDept(d.Ids)
}
