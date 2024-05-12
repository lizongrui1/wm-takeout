package admin

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/global"
	"wm-take-out/internal/api/controller"
	"wm-take-out/internal/middleware"
	"wm-take-out/internal/repository/dao"
	"wm-take-out/internal/service"
)

type CategoryRouter struct {
}

func (cr *CategoryRouter) RouterInit(group *gin.RouterGroup) {
	privateRouter := group.Group("category")
	privateRouter.Use(middleware.VerifyJWTAdmin())
	categoryCtrl := controller.NewCategoryController(service.NewCategoryService(dao.NewCategoryDao(global.DB)))
	{
		privateRouter.PUT("", categoryCtrl.UpdateCategory)
		privateRouter.GET("page", categoryCtrl.PageQuery)
		privateRouter.POST("status/:status", categoryCtrl.SetStatus)
		privateRouter.POST("", categoryCtrl.AddCategory)
		privateRouter.DELETE("", categoryCtrl.DeleteCategory)
		privateRouter.GET("list", categoryCtrl.List)
	}
}
