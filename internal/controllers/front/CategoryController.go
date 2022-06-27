package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/service/cate_service"
	"shop/pkg/app"
	"shop/pkg/constant"
)

// category api
type CategoryController struct {
}

// @Title 获取树形数据
// @Description 获取树形数据
// @Success 200 {object} app.Response
// @router /api/v1/category [get]
func (e *CategoryController) GetCateList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	cateService := cate_service.Cate{Enabled: 1}
	vo := cateService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)

}
