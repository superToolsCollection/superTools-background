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
* @Date: 2021-02-14 20:48
* @Description:
**/

type SpGoodController struct {
	GoodService service.ISpGoodService
}

func NewSpGoodController(goodService service.ISpGoodService) SpGoodController {
	return SpGoodController{GoodService: goodService}
}

func (s SpGoodController) GetGoodList(c *gin.Context) {
	param := service.GetGoodsListRequest{}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, totalPage, err := s.GoodService.GetGoodList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "SpGoodService.GetGoodList errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorGetGoodListFail)
		return
	}
	data := gin.H{
		"total": totalPage,
		"page":  pager.Page,
		"goods": result,
	}
	response.ToResponse(data, "获取成功", http.StatusOK)
}
