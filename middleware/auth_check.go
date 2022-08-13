package middleware

import (
	"gitee.com/phper95/pkg/sign"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"shop/internal/models"
	"shop/internal/service/auth_service"
	"shop/pkg/app"
	"shop/pkg/constant"
	"shop/pkg/global"
	"shop/pkg/jwt"
	"shop/pkg/logging"
	"shop/pkg/runtime"
	"strings"
)

const bearerLength = len("Bearer ")

func AppJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var appG = app.Gin{C: c}

		mytoken := c.Request.Header.Get("Authorization")
		if len(mytoken) < bearerLength {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			c.Abort()
			return
		}
		token := strings.TrimSpace(mytoken[bearerLength:])
		usr, err := jwt.ValidateToken(token)
		if err != nil {
			global.LOG.Error(err)
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH_CHECK_TOKEN_FAIL, data)
			c.Abort()
			return
		}

		c.Set(constant.AppAuthUser, usr)

		c.Next()

	}
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var appG = app.Gin{C: c}

		url := c.Request.URL.Path

		method := strings.ToLower(c.Request.Method)
		//部署线上开启
		//prohibit := "post,put,delete"
		//if url != "/admin/auth/logout" && strings.Contains(prohibit,method) {
		//	ctx.Output.JSON(controllers.ErrMsg("演示环境禁止操作",40006),
		//		true,true)
		//	return
		//}

		mytoken := c.Request.Header.Get("Authorization")
		if len(mytoken) < bearerLength {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			c.Abort()
			return
		}
		token := strings.TrimSpace(mytoken[bearerLength:])
		usr, err := jwt.ValidateToken(token)
		if err != nil {
			logging.Info(err)
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH_CHECK_TOKEN_FAIL, data)
			c.Abort()
			return
		}

		c.Set(constant.ContextKeyUserObj, usr)
		//url排除
		if urlExclude(url) {
			c.Next()
			return
		}

		//casbin check
		cb := runtime.Runtime.GetCasbinKey(constant.CASBIN)

		for _, roleName := range usr.Roles {
			//超级管理员过滤掉
			if roleName == "admin" {
				break
			}
			logging.Info(roleName, url, method)
			res, err := cb.Enforce(roleName, url, method)
			if err != nil {
				logging.Error(err)
			}
			//logging.Info(res)

			if !res {
				appG.Response(http.StatusForbidden, constant.ERROR_AUTH_CHECK_FAIL, data)
				c.Abort()
				return
			}
		}

		c.Next()

	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		appG := &app.Gin{C: c}
		// header信息校验
		authorization := c.GetHeader(constant.HeaderAuthField)
		authorizationDate := c.GetHeader(constant.HeaderAuthDateField)
		if len(authorization) == 0 || len(authorizationDate) == 0 {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error(" empty authorization header info", authorization, authorizationDate)
			c.Abort()
			return
		}
		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error("authorizationSplit error", authorizationSplit)
			c.Abort()
			return
		}

		//解析参数
		err := c.Request.ParseForm()
		if err != nil {
			appG.Response(http.StatusForbidden, constant.INVALID_PARAMS, data)
			c.Abort()
			return
		}
		key := authorizationSplit[0]
		authService := auth_service.Auth{}
		auth, err := authService.DetailByKey(key)
		if err != nil {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error("DetailByKey error", err, authorizationSplit)
			c.Abort()
			return
		}

		if auth.IsUsed == models.IsUsedNo {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error("IsUsed error", authorizationSplit)
			c.Abort()
			return
		}

		ok, err := sign.New(key, auth.BusinessSecret, constant.HeaderSignTokenTimeoutSeconds).Verify(authorization, authorizationDate,
			c.Request.URL.Path, c.Request.Method, c.Request.Form)
		if err != nil {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error("sign verify error", err)
		}
		if !ok {
			appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, data)
			logging.Error("sign verify not ok")
			c.Abort()
			return
		}
		c.Next()
	}

}

//url排除
func urlExclude(url string) bool {
	//公共路由直接放行
	reg := regexp.MustCompile(`[0-9]+`)
	newUrl := reg.ReplaceAllString(url, "*")
	var ignoreUrls = "/admin/menu/build,/admin/user/center,/admin/user/updatePass,/admin/auth/info," +
		"/admin/auth/logout,/admin/materialgroup/*,/admin/material/*,/shop/product/isFormatAttr/*"
	if strings.Contains(ignoreUrls, newUrl) {
		return true
	}

	return false
}
