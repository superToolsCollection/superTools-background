package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"superTools-background/pkg/app"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-09-23 20:27
* @Description: 用于配置jwt鉴权中间件，根据客户端传递过来的token坚定用户是否有权利访问加下来的接口
**/

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			id    string
			ecode = errcode.Success
			user *app.User
			err error
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if s, exist := c.GetQuery("x-user-id"); exist {
			id = s
		} else {
			id = c.GetHeader("x-user-id")
		}
		if token == "" || id == "" {
			ecode = errcode.InvalidParams
		} else {
			user, err = app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				userId, _ := strconv.Atoi(id)
				if userId != user.UserId {
					ecode = errcode.UnauthorizedTokenError
				} else {
					ad := &app.AccessDetails{
						AccessUuid: user.AccessUuid,
						UserId:     userId,
					}
					id, AuthErr := app.GetAuth(ad)
					if AuthErr != nil {
						ecode = errcode.UnauthorizedTokenError
					} else if id != user.UserId {
						ecode = errcode.UnauthorizedTokenError
					}
				}
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Set("accessUuid", user.AccessUuid)
		c.Next()
	}
}
