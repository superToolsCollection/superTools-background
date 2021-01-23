package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-23 20:09
* @Description: 在auth表内获取appKey以及appSecret
**/
type Auth struct {
	ID string `json:"id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type IAuth interface {
	GetAuth(appKey, appSecret string) (*Auth, error)
}

type AuthManager struct {
	table string
	conn  *gorm.DB
}

func (m *AuthManager)GetAuth(appKey, appSecret string) (*Auth, error){
	auth := &model.Auth{}
	result := m.conn.Where("app_key = ? AND app_secret = ? AND is_del = ?", appKey, appSecret, 0).Find(auth)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return &Auth{
		ID:auth.ID,
		AppKey:auth.AppKey,
		AppSecret:auth.AppSecret,
	}, nil
}

func NewAuthManager(table string, conn *gorm.DB) IAuth{
	return &AuthManager{table:table, conn:conn}
}