package server

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/internal/handler"
	"wm-take-out/internal/middleware"
	"wm-take-out/pkg/helper/resp"
	"wm-take-out/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	return r
}
