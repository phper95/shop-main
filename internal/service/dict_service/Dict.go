package dict_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
)

type Dict struct {
	Id      int64
	Name    string
	Enabled int

	PageNum  int
	PageSize int

	M *models.SysDict

	Ids []int64
}

func (d *Dict) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Enabled >= 0 {
		maps["enabled"] = d.Enabled
	}
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllDict(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Dict) Insert() error {
	return models.AddDict(d.M)
}

func (d *Dict) Save() error {
	return models.UpdateByDict(d.M)
}

func (d *Dict) Del() error {
	return models.DelByDict(d.Ids)
}
