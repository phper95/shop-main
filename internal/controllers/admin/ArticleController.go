package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"shop/internal/models"
	"shop/internal/service/article_service"
	"shop/pkg/app"
	"shop/pkg/constant"
	"shop/pkg/global"
	"shop/pkg/util"
)

// 文章api
type ArticleController struct {
}

// @Title 文章
// @Description 文章
// @Success 200 {object} app.Response
// @router / [get]
func (e *ArticleController) Get(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	global.LOG.Info(id)
	articleService := article_service.Article{
		Id: id,
	}
	vo := articleService.Get()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 文章列表
// @Description 文章列表
// @Success 200 {object} app.Response
// @router / [get]
func (e *ArticleController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	articleService := article_service.Article{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := articleService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 文章添加
// @Description 文章添加
// @Success 200 {object} app.Response
// @router / [post]
func (e *ArticleController) Post(c *gin.Context) {
	var (
		model models.WechatArticle
		appG  = app.Gin{C: c}
	)

	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	articleService := article_service.Article{
		M: &model,
	}

	if err := articleService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title 文章修改
// @Description 文章修改
// @Success 200 {object} app.Response
// @router / [put]
func (e *ArticleController) Put(c *gin.Context) {
	var (
		model models.WechatArticle
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	articleService := article_service.Article{
		M: &model,
	}

	if err := articleService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 文章删除
// @Description 文章删除
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *ArticleController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	articleService := article_service.Article{Ids: ids}

	if err := articleService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 发布文章
// @Description 发布文章
// @Success 200 {object} app.Response
// @router / [get]
func (e *ArticleController) Pub(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	global.LOG.Info(id)
	articleService := article_service.Article{
		Id: id,
	}
	if err := articleService.Pub(); err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
