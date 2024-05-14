package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/global"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/service"
)

type SetmealController struct {
	service service.SetMealService
}

func NewSetmealController(service service.SetMealService) *SetmealController {
	return &SetmealController{service: service}
}

func (sc *SetmealController) EditSetMeal(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.SetMealDTO
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param SetMealDTO json failed", "Err:", err.Error())
		return
	}
	err = sc.service.EditSetMeal(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EditSetMeal Error", "Err", err.Error())
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

func (sc *SetmealController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.SetMealPageQueryDTO
	dto.Name = ctx.Query("name")
	dto.Page, _ = strconv.Atoi(ctx.Query("page"))
	dto.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	dto.Status, _ = strconv.Atoi(ctx.Query("status"))
	result, err := sc.service.PageQuery(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param PageQuery json failed", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: result,
	})
}

func (sc *SetmealController) SetStatus(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	err := sc.service.SetStatus(ctx, id, status)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("SetStatus Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (sc *SetmealController) GetById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	result, err := sc.service.GetById(ctx, id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("GetById Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: result,
	})
}

func (sc *SetmealController) AddSetmeal(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.SetMealDTO
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
	err = sc.service.AddSetmeal(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Add Setmeal Failed", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (sc *SetmealController) DeleteSetmeal(ctx *gin.Context) {
	code := e.SUCCESS
	ids := ctx.Query("ids")
	err := sc.service.DeleteSetmeal(ctx, ids)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("delete setmeal error", "Err:", err.Error())
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
