package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2021-01-28 19:07
* @Description:
**/

type SpPermissionController struct {
	SpPermissionService service.ISpPermissionService
}

func NewSpPermissionController(spPermissionService service.ISpPermissionService) SpPermissionController {
	return SpPermissionController{SpPermissionService: spPermissionService}
}

func (s SpPermissionController) GetRights(c *gin.Context) {
	response := app.NewResponse(c)
	typeStr := strings.TrimSpace(c.Param("type"))
	if typeStr == "" || len(typeStr) == 0 || (typeStr != "list" && typeStr != "tree"){
		global.Logger.Errorf(c, "SpPermissionController.GetRights errs: %v", errors.New("wrong type"))
		response.ToErrorResponse(errcode.ErrorGetRightsFail)
		return
	}
	param := service.GetRightsRequest{
		Type: typeStr,
	}
	result, err := s.SpPermissionService.GetRights(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpPermissionService.GetRights errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRightsFail)
		return
	}
	response.ToResponse(result, "获取权限列表成功", http.StatusOK)
}
