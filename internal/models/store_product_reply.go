package models

import (
	"shop/pkg/global"
	"time"
)

type StoreProductReply struct {
	Uid                  int64     `json:"uid"`
	ProductId            int64     `json:"product_id"`
	Oid                  int64     `json:"oid"`
	Unique               string    `json:"unique"`
	ReplyType            string    `json:"reply_type"`
	ProductScore         int       `json:"product_score"`
	ServiceScore         int       `json:"service_score"`
	Comment              string    `json:"comment"`
	Pics                 string    `json:"pics"`
	MerchantReplyContent string    `json:"merchant_reply_content"`
	MerchantReplyTime    time.Time `json:"merchant_reply_time"`
	IsReply              int8      `json:"is_reply"`
	BaseModel
}

func (StoreProductReply) TableName() string {
	return "store_product_reply"
}

// get all
func GetAllProductReply(pageNUm int, pageSize int, maps interface{}) (int64, []StoreProductReply) {
	var (
		total int64
		data  []StoreProductReply
	)
	global.Db.Model(&StoreProductReply{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddStoreProductReply(m *StoreProductReply) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreProductReply(m *StoreProductReply) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreProductReply(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&StoreProductReply{}).Error
	if err != nil {
		return err
	}

	return err
}
