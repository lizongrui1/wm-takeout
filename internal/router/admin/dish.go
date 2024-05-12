package admin

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/global"
	"wm-take-out/internal/api/controller"
	"wm-take-out/internal/middleware"
	"wm-take-out/internal/repository/dao"
	"wm-take-out/internal/service"
)

type DishRouter struct {
}

func (dr *DishRouter) RouterInit(group *gin.RouterGroup) {
	privateRouter := group.Group("dish")
	privateRouter.Use(middleware.VerifyJWTAdmin())
	dishCtrl := controller.NewDishController(service.NewDishSe(dao.NewDishRepo(global.DB), dao.NewDishFlavorDao()))
	privateRouter.PUT("", dishCtrl.UpdateDish)
	privateRouter.DELETE("", dishCtrl.DeleteDish)
	privateRouter.POST("", dishCtrl.InsertDish)
	privateRouter.GET(":id", dishCtrl.GetDishById)
	privateRouter.GET("list", dishCtrl.List)
	privateRouter.GET("page", dishCtrl.PageQuery)
	privateRouter.POST("status/:status", dishCtrl.SetStatus)
}
