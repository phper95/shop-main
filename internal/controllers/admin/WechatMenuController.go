package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/service/wechat_menu_service"
	dto2 "shop/internal/service/wechat_menu_service/dto"
	"shop/pkg/app"
	"shop/pkg/constant"
)

// 菜单api
type WechatMenuController struct {
}

// @Title 获取菜单
// @Description 获取菜单
// @Success 200 {object} app.Response
// @router / [get]
func (e *WechatMenuController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	meuService := wechat_menu_service.Menu{}
	vo := meuService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 菜单更新
// @Description 菜单更新
// @Success 200 {object} app.Response
// @router / [post]
func (e *WechatMenuController) Post(c *gin.Context) {
	var (
		dto  dto2.WechatMenu
		appG = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	meuService := wechat_menu_service.Menu{
		Dto: dto,
	}

	if err := meuService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}
