package service

import (
	"errors"
	"superTools-background/internal/dao"
)

/**
* @Author: super
* @Date: 2020-09-23 20:09
* @Description: 用于Auth入参验证与service代码
**/

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

type Auth struct {
	ID string `json:"id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type IAuthService interface {
	CheckAuth(param *AuthRequest) error
}

type AuthService struct {
	authDao dao.IAuth
}

func NewAuthService(authDao dao.IAuth) IAuthService {
	return &AuthService{authDao: authDao}
}

func (svc *AuthService) CheckAuth(param *AuthRequest) error {
	auth, err := svc.authDao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}

	if auth.ID != "" {
		return nil
	}

	return errors.New("auth info does not exist")
}
