package address_service

import (
	"encoding/json"
	"errors"
	"gitee.com/phper95/pkg/cache"
	"shop/internal/models"
	"shop/internal/models/vo"
	"shop/internal/params"
	"shop/pkg/constant"
	"shop/pkg/global"
	"shop/pkg/util"
)

type Address struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.UserAddress

	Ids []int64

	Param *params.AddressParan
	Uid   int64
}

// del地址
func (d *Address) DelAddress() error {
	err := global.Db.
		Where("uid = ?", d.Uid).
		Where("id = ?", d.Id).
		Delete(&models.UserAddress{}).Error
	if err != nil {
		global.LOG.Error(err)
		return errors.New("操作失败")
	}

	return nil
}

// 设置默认地址
func (d *Address) SetDefault() error {
	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Model(&models.UserAddress{}).
		Where("uid = ?", d.Uid).Update("is_default", 0).Error
	if err != nil {
		global.LOG.Error(err)
		return errors.New("操作失败")
	}
	err = tx.Model(&models.UserAddress{}).
		Where("id = ?", d.Id).Update("is_default", 1).Error
	if err != nil {
		global.LOG.Error(err)
		return errors.New("操作失败")
	}
	return nil
}

//get list
func (d *Address) GetList() ([]models.UserAddress, int, int) {
	maps := make(map[string]interface{})
	maps["uid"] = d.Uid
	total, list := models.GetAllUserAddress(d.PageNum, d.PageSize, maps)

	totalNum := util.Int64ToInt(total)
	totalPage := util.GetTotalPage(totalNum, d.PageSize)
	return list, totalNum, totalPage
}

//add or update
func (d *Address) AddOrUpdate() (int64, error) {
	var err error
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	userAddress := &models.UserAddress{
		City:     d.Param.Address.City,
		CityId:   d.Param.Address.CityId,
		District: d.Param.Address.District,
		Province: d.Param.Address.Province,
		Detail:   d.Param.Detail,
		Uid:      d.Uid,
		Phone:    d.Param.Phone,
		PostCode: d.Param.PostCode,
		RealName: d.Param.RealName,
	}
	if d.Param.IsDefault {
		userAddress.IsDefault = 1
		err = tx.Model(&models.UserAddress{}).
			Where("uid = ?", d.Uid).Update("is_default", 0).Error
		if err != nil {
			global.LOG.Error(err)
			return 0, errors.New("操作失败")
		}
	}
	if d.Param.Id == 0 {
		err = tx.Create(userAddress).Error
		if err != nil {
			global.LOG.Error(err)
			return 0, errors.New("操作失败")
		}
	} else {
		err = tx.Model(&models.UserAddress{}).
			Where("id = ?", d.Param.Id).
			Updates(userAddress).Error
		if err != nil {
			global.LOG.Error(err)
			return 0, errors.New("操作失败")
		}
	}
	return userAddress.Id, nil
}

//get city list
func (d *Address) GetCities() []models.SystemCity {
	key := constant.CityList
	val, err := cache.GetRedisClient(cache.DefaultRedisClient).GetStr(key)
	if err != nil {
		global.LOG.Error("redis error ", err, "key", key, "cmd : Get", "client", cache.DefaultRedisClient)
	} else {
		var city []models.SystemCity
		err = json.Unmarshal([]byte(val), &city)
		if err != nil {
			global.LOG.Error(" json.Unmarshal error val : ", val)
		}
		if len(city) > 0 {
			return city
		}
	}
	maps := make(map[string]interface{})
	maps["is_show"] = 1
	list := models.GetAllSystemCity(maps)
	listCache, _ := json.Marshal(list)
	err = cache.GetRedisClient(cache.DefaultRedisClient).Set(key, listCache, 0)
	global.LOG.Error("set Cities data error", err, "key", key)
	return list
}

func (d *Address) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllUserAddress(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Address) Insert() error {
	return models.AddUserAddress(d.M)
}

func (d *Address) Save() error {
	return models.UpdateByUserAddress(d.M)
}

func (d *Address) Del() error {
	return models.DelByUserAddress(d.Ids)
}
