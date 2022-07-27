package front

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"shop/internal/params"
	"shop/internal/service/product_relation_service"
	"shop/internal/service/product_reply_service"
	"shop/internal/service/product_service"
	"shop/pkg/app"
	"shop/pkg/constant"
	productEnum "shop/pkg/enums/product"
	"shop/pkg/jwt"
	"shop/pkg/util"
)

// product api
type ProductController struct {
}

// @Title 获取商品列表数据
// @Description 获取商品列表数据
// @Success 200 {object} app.Response
// @router /api/v1/products [get]
func (e *ProductController) GoodsList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	productService := product_service.Product{
		Name:       c.Query("keyword"),
		Enabled:    1,
		PageNum:    util.GetFrontPage(c),
		PageSize:   util.GetFrontLimit(c),
		Sid:        c.Query("sid"),
		News:       c.Query("news"),
		PriceOrder: c.Query("price_order"),
		SalesOrder: c.Query("sales_order"),
	}

	vo, total, page := productService.GetList()

	appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)

}

// @Title 商品搜索接口
// @Description 根据关键词搜索商品
// @Success 200 {object} app.Response
// @router /api/v1/product/search [get]
func (e *ProductController) GoodsSearch(c *gin.Context) {
	var appG = app.Gin{C: c}
	productService := product_service.Product{
		Name:       c.Query("keyword"),
		Enabled:    1,
		PageNum:    util.GetFrontPage(c),
		PageSize:   util.GetFrontLimit(c),
		Sid:        c.Query("sid"),
		News:       c.Query("news"),
		PriceOrder: c.Query("price_order"),
		SalesOrder: c.Query("sales_order"),
	}
	if len(productService.Name) == 0 {
		vo, total, page := productService.GetList()
		appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)
		return
	}
	userid, _ := jwt.GetAppUserId(c)
	productService.Uid = userid
	vo, total, page := productService.SearchGoods()
	if vo == nil {
		appG.ResponsePage(http.StatusInternalServerError, constant.ERROR, nil, total, page)
		return
	}
	appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)
	return
}

// @Title 获取推荐商品
// @Description 获取推荐商品
// @Success 200 {object} app.Response
// @router /api/v1/product/hot [get]
func (e *ProductController) GoodsRecommendList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	productService := product_service.Product{
		Enabled:  1,
		PageNum:  0,
		PageSize: 6,
		Order:    productEnum.STATUS_1,
	}

	vo, _, _ := productService.GetList()

	appG.Response(http.StatusOK, constant.SUCCESS, vo)

}

// @Title 获取商品详情
// @Description 获取商品详情
// @Success 200 {object} app.Response
// @router /api/v1/product/detail/:id [get]
func (e *ProductController) GoodDetail(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		uid  int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	user, err := jwt.GetAppDetailUser(c)
	if err != nil {
		uid = 0
	} else {
		uid = user.Id
	}

	productService := product_service.Product{
		Id:  id,
		Uid: uid,
	}

	vo, err := productService.GetDetail()
	if err != nil {
		appG.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 添加收藏
// @Description 添加收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/add [post]
func (e *ProductController) AddCollect(c *gin.Context) {
	var (
		param params.RelationParam
		appG  = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		Param: &param,
		Uid:   uid,
	}
	if err := relationService.AddRelation(); err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, "success")

}

// @Title 取消收藏
// @Description 取消收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/del [post]
func (e *ProductController) DelCollect(c *gin.Context) {
	var (
		param params.RelationParam
		appG  = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		Param: &param,
		Uid:   uid,
	}
	if err := relationService.DelRelation(); err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, "success")

}

// @Title 获取商品评论列表数据
// @Description 获取商品评论列表数据
// @Success 200 {object} app.Response
// @router /api/v1/reply/list/:id [get]
func (e *ProductController) ReplyList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	replyService := product_reply_service.Reply{
		ProductId: com.StrTo(c.Param("id")).MustInt64(),
		PageNum:   util.GetFrontPage(c),
		PageSize:  util.GetFrontLimit(c),
		Type:      com.StrTo(c.Query("type")).MustInt(),
	}

	vo, total, page := replyService.GetList()

	appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)
}
