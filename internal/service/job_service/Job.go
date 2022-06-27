package job_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
)

type Job struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.SysJob

	Ids []int64
}

func (d *Job) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllJob(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Job) Insert() error {
	return models.AddJob(d.M)
}

func (d *Job) Save() error {
	return models.UpdateByJob(d.M)
}

func (d *Job) Del() error {
	return models.DelByJob(d.Ids)
}
