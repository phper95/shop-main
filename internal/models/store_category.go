package models

import "shop/pkg/global"

type StoreCategory struct {
	CateName string          `json:"cate_name" valid:"Required;"`
	Pid      int64           `json:"pid"`
	Sort     int             `json:"sort"`
	Pic      string          `json:"pic"`
	IsShow   int8            `json:"is_show"`
	Children []StoreCategory `gorm:"-" json:"children"`
	Label    string          `gorm:"-" json:"label"`
	BaseModel
}

func (StoreCategory) TableName() string {
	return "store_category"
}

func GetAllCates(maps interface{}) []StoreCategory {
	var data []StoreCategory
	global.Db.Where(maps).Find(&data)
	return RecursionCateList(data, 0)
}

//递归函数
func RecursionCateList(data []StoreCategory, pid int64) []StoreCategory {
	var listTree = make([]StoreCategory, 0)
	for _, value := range data {
		value.Label = value.CateName
		if value.Pid == pid {
			value.Children = RecursionCateList(data, value.Id)
			listTree = append(listTree, value)
		}
	}
	return listTree
}

func AddCate(m *StoreCategory) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByCate(m *StoreCategory) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByCate(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&StoreCategory{}).Error
	if err != nil {
		return err
	}

	return err
}
