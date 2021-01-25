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
}

func (s SpManagerController) AddUser(c *gin.Context) {
	param := service.AddSpMangerRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, err := s.SpManagerService.AddSpManager(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpManagerService.AddSpManager errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorAddUser)
		return
	}
	response.ToResponse(result, "用户创建成功", http.StatusCreated)
}

func (s SpManagerController) UpdateUserState(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpManagerService.UpdateUserState errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorUpdateUserStateFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	t := strings.TrimSpace(c.Param("type"))
	if t == "" || len(t) == 0 {
		global.Logger.Errorf(c, "SpManagerService.UpdateUserState errs: %v", errors.New("wrong type"))
		response.ToErrorResponse(errcode.ErrorUpdateUserStateFail)
		return
	}
	param := service.UpdateSpMangerStateRequest{
		ID:   id,
		Type: t,
	}
	result, err := s.SpManagerService.UpdateSpManagerState(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpManagerService.UpdateUserState errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserStateFail)
		return
	}
	data := gin.H{
		"id":       result.MgID,
		"rid":      result.RoleID,
		"username": result.MgName,
		"mobile":   result.MgMobile,
		"email":    result.MgEmail,
		"mg_state": result.MgState,
	}
	response.ToResponse(data, "设置状态成功", http.StatusOK)
}

func (s SpManagerController) GetUserByID(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpManagerService.GetUserByID errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorGetUserByID)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param := service.GetSpMangerByIDRequest{
		ID:int(id),
	}
	result, err := s.SpManagerService.GetSpManagerByID(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpManagerService.UpdateUserState errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserByID)
		return
	}
	data := gin.H{
		"id":       result.MgID,
		"rid":      result.RoleID,
		"username": result.MgName,
		"mobile":   result.MgMobile,
		"email":    result.MgEmail,
	}
	response.ToResponse(data, "查询成功", http.StatusOK)
}
