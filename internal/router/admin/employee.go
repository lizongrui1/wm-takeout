package admin

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/global"
	"wm-take-out/internal/api/controller"
	"wm-take-out/internal/middleware"
	"wm-take-out/internal/repository/dao"
	"wm-take-out/internal/service"
)

type EmployeeRouter struct {
	service service.EmployeeService
}

func (er *EmployeeRouter) RouterInit(group *gin.RouterGroup) {
	publicRouter := group.Group("employee")  // 公开接口
	privateRouter := group.Group("employee") // 管理员接口
	privateRouter.Use(middleware.VerifyJWTAdmin())
	er.service = service.NewEmployeeService(dao.NewEmployeeDao(global.DB))
	employeeCtl := controller.NewEmployeeController(er.service)
	{
		publicRouter.POST("/login", employeeCtl.Login)
		privateRouter.POST("/logout", employeeCtl.Logout)
		privateRouter.PUT("/editPassword", employeeCtl.UpdatePassword)
		privateRouter.POST("/status/:status", employeeCtl.EmployeeStatus)
		privateRouter.GET("/page", employeeCtl.PageQuery)
		privateRouter.POST("", employeeCtl.AddEmployee)
		privateRouter.GET(":id", employeeCtl.EmployeeQueryById)
		privateRouter.PUT("", employeeCtl.UpdateEmployee)
	}
}
