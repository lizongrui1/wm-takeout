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

type DishController struct {
	service service.DishService
}

func (dc *DishController) NewDishController(service service.DishService) *DishController {
	return &DishController{
		service: service,
	}
}

func (dc *DishController) InsertDish(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.DishDTO
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("Dish Binding Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	err = dc.service.InsertDish(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("InsertDish Error", "Err:", err.Error())
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

func (dc *DishController) DeleteDish(ctx *gin.Context) {
	code := e.SUCCESS
	ids := ctx.Query("ids")
	err := dc.service.DeleteDish(ctx, ids)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("DeleteDish Error", "Err:", err.Error())
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

func (dc *DishController) UpdateDish(ctx *gin.Context) {
	code := e.SUCCESS
	dto := request.DishUpdateDTO{}
	err := dc.service.UpdateDish(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Param UpdateDish Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	err := dc.service.UpdateDish(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("UpdateDish Error", "Err:", err.Error())
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

func (dc *DishController) GetDishById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	result, err := dc.service.GetDishById(ctx, id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("GetDishById Error", "Err:", err.Error())
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

func (dc *DishController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.DishPageQueryDTO
	dto.Name = ctx.Query("name")
	dto.Page, _ = strconv.Atoi(ctx.Query("page"))
	dto.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	dto.Status, _ = strconv.Atoi(ctx.Query("status"))
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("PageQuery Binding Error", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	result, _ := dc.service.PageQuery(ctx, &dto)
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: result,
	})
}
