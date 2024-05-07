package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/global"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func (ec *EmployeeController) Login(ctx *gin.Context) {
	code := e.SUCCESS
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("EmployeeController login 解析失败")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeController login Error:", err.Error())
		ctx.JSON(http.StatusUnauthorized, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (ec *EmployeeController) Logout(ctx *gin.Context) {
	code := e.SUCCESS
	err := ec.service.Logout(ctx)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeController login Error:", err.Error())
		ctx.JSON(http.StatusUnauthorized, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.EmployeeDTO
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("AddEmployee Error:", err.Error())
		ctx.JSON(http.StatusBadRequest, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	err = ec.service.AddEmployee(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddEmployee Error:", err.Error())
		ctx.JSON(http.StatusInternalServerError, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (ec *EmployeeController) UpdatePassword(ctx *gin.Context) {
	code := e.SUCCESS

}
