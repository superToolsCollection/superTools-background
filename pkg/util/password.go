package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

/**
* @Author: super
* @Date: 2021-01-23 20:54
* @Description:
**/

//将明文密码加密
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

}

//验证登录密码是否正确
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码校验错误")
	}
	return true, nil
}