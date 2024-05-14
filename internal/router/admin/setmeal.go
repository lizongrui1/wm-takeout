package admin

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/global"
	"wm-take-out/internal/api/controller"
	"wm-take-out/internal/middleware"
	"wm-take-out/internal/repository/dao"
	"wm-take-out/internal/service"
)

type SetmealRouter struct {
}

func (sr *SetmealRouter) RouterInit(group *gin.RouterGroup) {
	privateRouter := group.Group("setmeal")
	privateRouter.Use(middleware.VerifyJWTAdmin())
	setmealCrtl := controller.NewSetmealController(service.NewSetMealService(dao.NewSetMealDao(global.DB), dao.NewSetmealDishDao()))
	privateRouter.PUT("", setmealCrtl.EditSetMeal)
	privateRouter.GET("/page", setmealCrtl.PageQuery)
	privateRouter.POST("status/:status", setmealCrtl.SetStatus)
	privateRouter.DELETE("", setmealCrtl.DeleteSetmeal)
	privateRouter.POST("", setmealCrtl.AddSetmeal)
	privateRouter.GET("/:id", setmealCrtl.GetById)
}
