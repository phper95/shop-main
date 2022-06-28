package models

import (
	"github.com/astaxie/beego/validation"
	"shop/pkg/global"
)

type WechatArticle struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Image     string `json:"image"`
	Synopsis  string `json:"synopsis"`
	Content   string `json:"content"`
	Visit     int    `json:"visit"`
	Sort      int    `json:"sort"`
	Url       string `json:"url"`
	Status    int    `json:"status"`
	ProductId int    `json:"product_id"`
	MediaId   string `json:"media_id"`
	IsPub     int    `json:"is_pub"`
	BaseModel
}

func (WechatArticle) TableName() string {
	return "wechat_article"
}

func (a *WechatArticle) Valid(v *validation.Validation) {
	if a.Title == "" {
		v.SetError("title", "标题不能为空")
	}
	if a.Author == "" {
		v.SetError("author", "作者不能为空")
	}
}

// get all
func GetAllWechatArticle(pageNUm int, pageSize int, maps interface{}) (int64, []WechatArticle) {
	var (
		total int64
		data  []WechatArticle
	)

	global.Db.Model(&WechatArticle{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddWechatArticle(m *WechatArticle) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatArticle(m *WechatArticle) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByWechatArticle(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&WechatArticle{}).Error
	if err != nil {
		return err
	}

	return err
}
