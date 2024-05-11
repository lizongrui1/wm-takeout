package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/common/utils"
	"wm-take-out/global"
	"wm-take-out/internal/enum"
)

func VerifyJWTAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := e.SUCCESS
		token := context.Request.Header.Get(global.Config.Jwt.Admin.Name)
		payLoad, err := utils.ParseToken(token, global.Config.Jwt.Admin.Secret)
		if err != nil {
			code = e.UNKNOW_IDENTITY
			context.JSON(http.StatusUnauthorized, common.Result{
				Code: code,
			})
			context.Abort()
			return
		}
		context.Set(enum.CurrentId, payLoad.UserId)
		context.Set(enum.CurrentName, payLoad.GrantScope)
		context.Next()
	}
}

func VerifyJWYUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := e.SUCCESS
		token := context.Request.Header.Get(global.Config.Jwt.User.Name)
		payLoad, err := utils.ParseToken(token, global.Config.Jwt.User.Secret)
		if err != nil {
			code = e.ERROR
			context.JSON(http.StatusUnauthorized, common.Result{
				Code: code,
			})
			context.Abort()
			return
		}
		context.Set(enum.CurrentId, payLoad.UserId)
		context.Set(enum.CurrentId, payLoad.GrantScope)
	}
}
