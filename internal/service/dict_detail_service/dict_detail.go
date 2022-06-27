package dict_detail_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
)

type DictDetail struct {
	Id       int64
	DictName string

	DictId  int64
	Enabled int

	PageNum  int
	PageSize int

	M *models.SysDictDetail

	Ids []int64
}

func (d *DictDetail) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.DictId >= 0 {
		maps["dict_id"] = d.DictId
	}

	if d.DictName != "" {
		maps["dict_name"] = d.DictName
	}

	total, list := models.GetAllDictDetail(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *DictDetail) Insert() error {
	return models.AddDictDetail(d.M)
}

func (d *DictDetail) Save() error {
	return models.UpdateByDictDetail(d.M)
}

func (d *DictDetail) Del() error {
	return models.DelByDictDetail(d.Ids)
}
