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

type CategoryController struct {
	service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		service: categoryService,
	}
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.CategoryDTO
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param CategoryDTO json failed", err.Error())
		return
	}
	err = cc.service.UpdateCategory(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("UpdateCategory Error", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.CategoryPageQueryDTO
	var err error
	dto.Name = ctx.Query("name")
	dto.Page, _ = strconv.Atoi(ctx.Query("page"))
	dto.PageSize, err = strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		dto.Page = 1 // 设置默认页码
	}
	query, err := cc.service.PageQuery(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("PageQuery Error", err.Error())
		// 这里到底要不要＋ctx.Json？
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: query,
	})
}

func (cc *CategoryController) SetStatus(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	err := cc.service.SetStatus(ctx, id, status)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("SetStatus Error", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) AddCategory(ctx *gin.Context) {
	code := e.SUCCESS
	var dto request.CategoryDTO
	err := ctx.Bind(&dto)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param CategoryDTO json failed", err.Error())
		return
	}
	err = cc.service.AddCategory(ctx, dto)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddCategory Error", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	err := cc.service.DeleteCategory(ctx, id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("DeleteCategory Error", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) List(ctx *gin.Context) {
	code := e.SUCCESS
	//var dto request.CategoryDTO
	cate, _ := strconv.Atoi(ctx.Query("type"))
	list, err := cc.service.List(ctx, cate)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("List Query Error", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: list,
	})
}
