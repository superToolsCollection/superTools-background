package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
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
