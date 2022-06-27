package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/service/log_service"
	"shop/pkg/app"
	"shop/pkg/constant"
	"shop/pkg/util"
)

// 角色 API
type LogController struct {
}

// @Title 日志列表
// @Description 日志列表
// @Success 200 {object} app.Response
// @router / [get]
func (e *LogController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	blurry := c.DefaultQuery("blurry", "")
	logService := log_service.Log{
		Des:      blurry,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := logService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 日志删除
// @Description 日志删除
// @Success 200 {object} app.Response
// @router / [delete]
func (e *LogController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	c.BindJSON(&ids)
	logService := log_service.Log{Ids: ids}

	if err := logService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
