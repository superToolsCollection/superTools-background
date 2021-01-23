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
* @Date: 2021-01-23 20:35
* @Description:
**/

type SpManagerController struct {
	SpManagerService service.ISpManagerService
}

func NewSpManagerController(spManagerService service.ISpManagerService) SpManagerController {
	return SpManagerController{SpManagerService: spManagerService}
}

func (s SpManagerController) Login(c *gin.Context) {
	param := service.GetSpMangerRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	spManager, err := s.SpManagerService.GetSpManager(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.SignIn err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserSignInFail)
		return
	}

	token, err := app.GenerateTokenByUserName(spManager.MgName)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	data := gin.H{
		"id":       spManager.MgID,
		"rid":      spManager.RoleID,
		"username": spManager.MgName,
		"mobile":   spManager.MgMobile,
		"email":    spManager.MgEmail,
		"token":    token,
	}
	response.ToResponse(data, "登陆成功", http.StatusOK)
	return
}

func (s SpManagerController) Users(c *gin.Context) {
	param := service.GetSpMangerListRequest{}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, totalPage, err := s.SpManagerService.GetSpManagerList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "SpManagerService.GetSpManagerList errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorUserListFail)
		return
	}
	data := gin.H{
		"totalpage": totalPage,
		"pagenum":   pager.Page,
		"users":     result,
	}
	response.ToResponse(data, "获取成功", http.StatusOK)
	return
}
