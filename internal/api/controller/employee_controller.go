package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/global"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: employeeService}
}

func (ec *EmployeeController) Login(ctx *gin.Context) {
	code := e.SUCCESS
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("EmployeeController login 解析失败", "Err:", err.Error())
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeController login Error:", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
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
		global.Log.Warn("EmployeeController login Error:", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
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
		global.Log.Debug("Bind Error:", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	err = ec.service.AddEmployee(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddEmployee Error:", "Err:", err.Error())
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
	var cp request.EmployeeChangePassword
	err := ctx.Bind(&cp)
	if err != nil {
		global.Log.Debug("Bind Error:", "Err:", err.Error())
		return
	}
	if id, ok := ctx.Get(enum.CurrentId); ok {
		cp.EmpId = id.(uint64)
	}
	err = ec.service.UpdatePassword(ctx, cp)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("UpdatePassword  Error:", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (ec *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	code := e.SUCCESS
	var upe request.EmployeeDTO
	err := ctx.Bind(&upe)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("Bind Error:", "Err:", err.Error())
		return
	}
	err = ec.service.UpdateEmployee(ctx, upe)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("UpdateEmployee Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (ec *EmployeeController) EmployeeQueryById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	employee, err := ec.service.EmployeeQueryById(ctx, id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeQueryById", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: employee,
	})
}

func (ec *EmployeeController) EmployeeStatus(ctx *gin.Context) {
	code := e.SUCCESS
	status, _ := strconv.Atoi(ctx.Query("status"))
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err := ec.service.EmployeeStatus(ctx, id, status)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeStatus Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	global.Log.Info("Employee Status：", "id", id, "status:", status)
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (ec *EmployeeController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var epq request.EmployeePageQueryDTO
	epq.Name = ctx.Query("name")
	epq.Page, _ = strconv.Atoi(ctx.Query("page"))
	epq.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	result, err := ec.service.PageQuery(ctx, epq)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("PageQuery Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: result,
	})
}
