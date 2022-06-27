package product_reply_service

import (
	"github.com/jinzhu/copier"
	"shop/internal/models"
	"shop/internal/models/vo"
	vo2 "shop/internal/service/product_service/vo"
	"shop/pkg/global"
	"shop/pkg/util"
)

type Reply struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.StoreProductReply

	Ids []int64

	Uid       int64
	ProductId int64
	Type      int
}

////add collect
//func (d *Relation) AddRelation() error {
//	//productId := com.StrTo(d.Param.Id).MustInt64()
//	if IsRelation(d.Param.Id,d.Uid) {
//		return errors.New("已经收藏过")
//	}
//	model := &models.StoreProductRelation{
//		Uid: d.Uid,
//		ProductId: d.Param.Id,
//		Type: d.Param.Category,
//	}
//	return models.AddStoreProductRelation(model)
//}

////del collect
//func (d *Relation) DelRelation() error {
//	if !IsRelation(d.Param.Id,d.Uid) {
//		return errors.New("已经取消过")
//	}
//	err := global.Db.
//		Where("uid = ?",d.Uid).
//		Where("product_id = ?",d.Param.Id).
//		Where("type = ?",relationEnum.COLLECT).
//		Delete(&models.StoreProductRelation{}).Error
//	if err != nil {
//		global.LOG.Error(err)
//		return errors.New("取消失败")
//	}
//	return nil
//}

////是否收藏
//func IsRelation(productId,uid int64) bool  {
//	var (
//		count int64
//		error error
//	)
//	error = global.Db.Model(&models.StoreProductRelation{}).
//		Where("uid = ?",uid).
//		Where("product_id = ?",productId).
//		Where("type = ?",relationEnum.COLLECT).
//		Count(&count).Error
//	if error != nil {
//		global.LOG.Error(error)
//		return false
//	}
//	if count == 0 {
//		return false
//	}
//
//	return true
//}

//评论列表
func (d *Reply) GetList() ([]vo2.ProductReply, int, int) {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}
	if d.ProductId > 0 {
		maps["product_id"] = d.ProductId
	}

	var replyVo []vo2.ProductReply

	total, list := models.GetAllProductReply(d.PageNum, d.PageSize, maps)
	e := copier.Copy(&replyVo, list)
	if e != nil {
		global.LOG.Error(e)
	}
	totalNum := util.Int64ToInt(total)
	totalPage := util.GetTotalPage(totalNum, d.PageSize)
	return replyVo, totalNum, totalPage
}

func (d *Reply) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}
	if d.ProductId > 0 {
		maps["product_id"] = d.ProductId
	}

	total, list := models.GetAllProductReply(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Reply) Insert() error {
	return models.AddStoreProductReply(d.M)
}

func (d *Reply) Save() error {
	return models.UpdateByStoreProductReply(d.M)
}

func (d *Reply) Del() error {
	return models.DelByStoreProductReply(d.Ids)
}
