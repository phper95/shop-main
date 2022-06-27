package canvas_service

import (
	"shop/internal/models"
	"shop/internal/models/vo"
	"shop/pkg/global"
)

type Canvas struct {
	Id       int64
	Terminal int

	Enabled int

	M *models.StoreCanvas

	Ids []int64
}

func (d *Canvas) Get() vo.ResultList {
	var data models.StoreCanvas
	err := global.Db.Model(&models.StoreCanvas{}).Where("terminal = ?", d.Terminal).First(&data).Error
	if err != nil {
		global.LOG.Error(err)
	}
	return vo.ResultList{Content: data, TotalElements: 0}
}

func (d *Canvas) Save() error {
	if d.M.Id == 0 {
		return models.AddCanvas(d.M)
	} else {
		data := &models.StoreCanvas{
			Name:     d.M.Name,
			Terminal: d.M.Terminal,
			Json:     d.M.Json,
		}
		return global.Db.Model(&models.StoreCanvas{}).Where("id = ?", d.M.Id).Updates(data).Error
	}

}

//func (d *Canvas) Save() error {
//	return models.UpdateByCanvas(d.M)
//}

func (d *Canvas) Del() error {
	return models.DelByCanvas(d.Ids)
}
