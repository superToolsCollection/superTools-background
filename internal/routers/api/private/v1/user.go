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
* @Date: 2021-01-23 19:04
* @Description:
**/

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{UserService: userService}
}

// @Summary 用户登录
// @Tags user
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} service.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/login [post]
func (u UserController) Login(c *gin.Context) {
	param := service.UserSignInRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, err := u.UserService.SignIn(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.SignIn err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserSignInFail)
		return
	}
	token, err := app.GenerateTokenByUserName(user.UserName)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	data := gin.H{
		"id":       user.ID,
		"username": user.UserName,
		"token":    token,
	}
	response.ToResponse(data, "登陆成功", http.StatusOK)
	return
}
