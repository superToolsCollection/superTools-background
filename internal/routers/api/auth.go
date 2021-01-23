package api

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
* @Date: 2020-09-23 20:10
* @Description: auth对应的restful api
**/

type AuthController struct {
	AuthService service.IAuthService
}

func NewAuthController(authService service.IAuthService) AuthController {
	return AuthController{AuthService: authService}
}

// @Summary 获得token
// @Produce json
// @Param app_key query string true "app_key"
// @Param app_secret query string true "app_secret"
// @Success 200 {string} string "token"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [get]
func (controller *AuthController) GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := controller.AuthService.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	}, "获取token成功", http.StatusOK)
}
