package admin

import (
	"encoding/json"
	"gitee.com/phper95/pkg/mq"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"shop/internal/models"
	"shop/internal/service/product_service"
	dto2 "shop/internal/service/product_service/dto"
	"shop/pkg/app"
	"shop/pkg/constant"
	"shop/pkg/enums/product"
	"shop/pkg/global"
	"shop/pkg/util"
	"strconv"
)

// 商品 api
type StoreProductController struct {
}

// @Title 商品列表
// @Description 商品列表
// @Success 200 {object} app.Response
// @router / [get]
func (e *StoreProductController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("is_show", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	productService := product_service.Product{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := productService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 获取商品信息
// @Description 获取商品信息
// @Success 200 {object} app.Response
// @router /info/:id [get]
func (e *StoreProductController) GetInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	productService := product_service.Product{
		Id: id,
	}
	vo := productService.GetProductInfo()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 商品添加
// @Description 商品添加
// @Success 200 {object} app.Response
// @router /addOrSave [post]
func (e *StoreProductController) Post(c *gin.Context) {
	var (
		dto  dto2.StoreProduct
		appG = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	productService := product_service.Product{
		Dto: dto,
	}
	model, err := productService.AddOrSaveProduct()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	//发送变更事件消息
	defer func() {
		operation := product.OperationCreate
		if dto.Id > 0 {
			operation = product.OperationUpdate
		}
		productMsg := models.ProductMsg{operation, &model}
		msg, _ := json.Marshal(productMsg)
		p, o, e := mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Send(
			&sarama.ProducerMessage{
				Topic: product.Topic,
				Key:   mq.KafkaMsgValueStrEncoder(strconv.FormatInt(dto.Id, 10)),
				Value: mq.KafkaMsgValueEncoder(msg),
			},
		)
		if e != nil {
			global.LOG.Error("send msg error", e, "partion:", p, "offset", o, "id", dto.Id)
		}
	}()

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title 商品上下架
// @Description 商品上下架
// @Success 200 {object} app.Response
// @router /onsale/:id [post]
func (e *StoreProductController) OnSale(c *gin.Context) {
	var (
		dto  dto2.OnSale
		appG = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &dto)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}
	id := com.StrTo(c.Param("id")).MustInt64()
	productService := product_service.Product{
		SaleDto: dto,
		Id:      id,
	}

	if err := productService.OnSaleByProduct(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	//发送变更事件消息
	defer func() {
		operation := product.OperationOnSale
		if dto.Status == 0 {
			operation = product.OperationUnSale
		}
		productInfo := models.GetProduct(id)

		productMsg := models.ProductMsg{operation, &productInfo}
		msg, _ := json.Marshal(productMsg)
		p, o, e := mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Send(
			&sarama.ProducerMessage{
				Topic: product.Topic,
				Key:   mq.KafkaMsgValueStrEncoder(strconv.FormatInt(id, 10)),
				Value: mq.KafkaMsgValueEncoder(msg),
			},
		)
		if e != nil {
			global.LOG.Error("send msg error", e, "partion:", p, "offset", o, "id", id)
		}
	}()

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title 商品删除
// @Description 商品删除
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *StoreProductController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	productInfo := models.GetProduct(id)
	if productInfo.Id == 0 {
		appG.Response(http.StatusNotFound, constant.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}
	productService := product_service.Product{Ids: ids}

	if err := productService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	//发送变更事件消息
	defer func() {
		operation := product.OperationDelete
		productMsg := models.ProductMsg{operation, &productInfo}
		msg, _ := json.Marshal(productMsg)
		p, o, e := mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Send(
			&sarama.ProducerMessage{
				Topic: product.Topic,
				Key:   mq.KafkaMsgValueStrEncoder(strconv.FormatInt(id, 10)),
				Value: mq.KafkaMsgValueEncoder(msg),
			},
		)
		if e != nil {
			global.LOG.Error("send msg error", e, "partion:", p, "offset", o, "id", id)
		}
	}()

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 商品sku生成
// @Description 商品sku生成
// @Success 200 {object} app.Response
// @router /isFormatAttr/:id [post]
func (e *StoreProductController) FormatAttr(c *gin.Context) {
	var (
		appG    = app.Gin{C: c}
		jsonObj map[string]interface{}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	c.BindJSON(&jsonObj)
	productService := product_service.Product{
		Id:      id,
		JsonObj: jsonObj,
	}
	vo := productService.PublicFormatAttr()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}
