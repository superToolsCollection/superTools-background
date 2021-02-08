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
* @Date: 2021-02-07 20:38
* @Description:
**/

type SpRoleController struct {
	Service service.ISpRoleService
}

func NewSpRoleController(service service.ISpRoleService) SpRoleController {
	return SpRoleController{Service: service}
}

func (s SpRoleController) GetRoleList(c *gin.Context) {
	response := app.NewResponse(c)
	result, err := s.Service.GetRoles()
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.GetRoleList errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRoleListFail)
		return
	}
	response.ToResponse(result, "获取成功", http.StatusOK)
}

func (s SpRoleController) AddRole(c *gin.Context) {
	param := service.AddRoleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, err := s.Service.AddRole(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.AddRole errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorAddRoleFail)
		return
	}
	response.ToResponse(result, "创建成功", http.StatusCreated)
}

func (s SpRoleController) GetRoleById(c *gin.Context) {
	param := service.GetRoleByIdRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.GetRoleById errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorGetRoleByIdFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.ID = int(id)
	result, err := s.Service.GetRoleByID(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.GetRoleById errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRoleByIdFail)
		return
	}
	response.ToResponse(result, "获取成功", http.StatusOK)
}

func (s SpRoleController) UpdateRole(c *gin.Context) {
	param := service.UpdateRoleRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.UpdateRole errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorUpdateRoleFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.ID = int(id)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, err := s.Service.UpdateRole(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.UpdateRole errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorUpdateRoleFail)
		return
	}
	response.ToResponse(result, "编辑成功", http.StatusOK)
}

func (s SpRoleController) DeleteRole(c *gin.Context) {
	param := service.DeleteRoleRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.DeleteRole errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorDeleteRoleFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.ID = int(id)
	err := s.Service.DeleteRole(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.DeleteRole errs: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteRoleFail)
		return
	}
	response.ToResponse(gin.H{}, "删除成功", http.StatusOK)
}

func (s SpRoleController) UpdateRights(c *gin.Context) {
	param := service.UpdateRightRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("roleId"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.UpdateRights errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorUpdateRightFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.RoleID = int(id)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := s.Service.UpdateRight(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.DeleteRole errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateRoleFail)
		return
	}
	response.ToResponse(gin.H{}, "更新成功", http.StatusOK)
}

func (s SpRoleController) DeleteRight(c *gin.Context) {
	param := service.DeleteRightRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.DeleteRight errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorDeleteRightFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.RoleID = int(id)

	rightIdStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpRoleService.DeleteRight errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorDeleteRightFail)
		return
	}
	rightId := convert.StrTo(rightIdStr).MustInt64()
	param.RightId = int(rightId)

	err := s.Service.DeleteRight(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.DeleteRight errs: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteRightFail)
		return
	}
	role, err := s.Service.GetRoleByID(&service.GetRoleByIdRequest{ID: int(id)})
	response.ToResponse(role.Children, "取消权限成功", http.StatusOK)
}
