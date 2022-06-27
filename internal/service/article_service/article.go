package article_service

import (
	"errors"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"shop/internal/models"
	"shop/internal/models/vo"
	articleEnum "shop/pkg/enums/article"
	"shop/pkg/global"
	"strings"
)

type Article struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.WechatArticle

	Ids []int64
}

func (d *Article) Get() vo.ResultList {
	var data models.WechatArticle
	err := global.Db.Model(&models.WechatArticle{}).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		global.LOG.Error(err)
	}
	return vo.ResultList{Content: data, TotalElements: 0}
}

func (d *Article) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllWechatArticle(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Article) Pub() error {
	var data models.WechatArticle
	err := global.Db.Model(&models.WechatArticle{}).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		global.LOG.Error(err)
	}
	if data.IsPub == articleEnum.IS_PUB_1 {
		return errors.New("已经发布过啦！")
	}
	official := global.WechatOfficial
	m := official.GetMaterial()
	ss := strings.Replace(data.Image, global.CONFIG.App.PrefixUrl+"/", global.CONFIG.App.RuntimeRootPath, 1)
	mediaId, url, err := m.AddMaterial(material.MediaTypeThumb, ss)
	if err != nil {
		global.LOG.Error(err)
		return err
	}
	global.LOG.Info(mediaId, url)
	art := &material.Article{
		Title:            data.Title,
		ThumbMediaID:     mediaId,
		ThumbURL:         url,
		Author:           data.Author,
		Digest:           data.Synopsis,
		ShowCoverPic:     1,
		Content:          data.Content,
		ContentSourceURL: "",
	}
	arts := []*material.Article{art}
	id, err := m.AddNews(arts)
	global.LOG.Info(id, err)
	if err != nil {
		global.LOG.Error(err)
		return err
	}

	data.MediaId = id
	data.IsPub = articleEnum.IS_PUB_1

	return models.UpdateByWechatArticle(&data)
}

func (d *Article) Insert() error {
	return models.AddWechatArticle(d.M)
}

func (d *Article) Save() error {
	return models.UpdateByWechatArticle(d.M)
}

func (d *Article) Del() error {
	return models.DelByWechatArticle(d.Ids)
}
