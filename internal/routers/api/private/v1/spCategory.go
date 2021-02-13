package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/convert"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2021-02-08 14:55
* @Description:
**/

type SpCategoryController struct {
	Service service.ISpCategoryService
}

func NewSpCategoryController(service service.ISpCategoryService) SpCategoryController {
	return SpCategoryController{Service: service}
}

func (s SpCategoryController) GetCategoriesList(c *gin.Context) {
	param := service.GetCategoriesListRequest{}
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	list, err := s.Service.GetCategoriesList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "SpCategoryService.GetCategoriesList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetCategoriesListFail)
		return
	}
	response.ToResponse(list, "获取成功", http.StatusOK)
}

func (s SpCategoryController) AddCategory(c *gin.Context) {
	param := service.AddCategoryRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	category, err := s.Service.AddCategory(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpCategoryService.AddCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorAddCategoryFail)
		return
	}
	response.ToResponse(category, "创建成功", http.StatusCreated)
}

func (s SpCategoryController) GetCategory(c *gin.Context) {
	param := service.GetCategoryByIdRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpCategoryService.GetCategory errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorGetCategoryFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	category, err := s.Service.GetCategory(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpCategoryService.GetCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetCategoryFail)
		return
	}
	response.ToResponse(category, "获取成功", http.StatusOK)
}

func (s SpCategoryController) UpdateCategory(c *gin.Context) {
	param := service.UpdateCategoryRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpCategoryService.UpdateCategory errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorUpdateCategoryFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	category, err := s.Service.UpdateCategory(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpCategoryService.UpdateCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateCategoryFail)
		return
	}
	response.ToResponse(category, "更新成功", http.StatusOK)
}

func (s SpCategoryController) DeleteCategory(c *gin.Context){
	param := service.DeleteCategoryRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpCategoryService.DeleteCategory errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorDeleteCategoryFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	err := s.Service.DeleteCategory(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpCategoryService.DeleteCategory err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteCategoryFail)
		return
	}
	response.ToResponse(gin.H{}, "删除成功", http.StatusOK)
}