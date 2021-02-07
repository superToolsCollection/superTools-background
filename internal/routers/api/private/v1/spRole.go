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
* @Date: 2021-02-07 20:38
* @Description:
**/

type SpRoleController struct {
	Service service.ISpRoleService
}

func NewSpRoleController(service service.ISpRoleService) SpRoleController{
	return SpRoleController{Service:service}
}

func (s SpRoleController) GetRoleList(c *gin.Context){
	response := app.NewResponse(c)
	result, err := s.Service.GetRoles()
	if err != nil {
		global.Logger.Errorf(c, "SpRoleService.GetRoleList errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRoleListFail)
		return
	}
	response.ToResponse(result, "获取成功", http.StatusOK)
}