package service

import (
	"errors"
	"superTools-background/pkg/util"

	"superTools-background/internal/dao"
	"superTools-background/pkg/idGenerator"
)

/**
* @Author: super
* @Date: 2020-11-24 14:33
* @Description:
**/
type UserSignInRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type UserRegisterRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	NickName string `form:"nick_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type UserUpdateInfoRequest struct {
	ID       string `form:"id" binding:"required,min=2,max=4294967295"`
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	NickName string `form:"nick_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type User struct {
	ID           string `json:"id"`
	NickName     string `json:"nick_name"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"-"`
}

type IUserService interface {
	SignIn(param *UserSignInRequest) (*User, error)
	Register(param *UserRegisterRequest) (string, error)
	UpdateInfo(param *UserUpdateInfoRequest) error
}

type UserService struct {
	userDao dao.IUser
}

func (s *UserService) SignIn(param *UserSignInRequest) (*User, error) {
	user, err := s.userDao.SelectByUserName(param.UserName)
	if err != nil {
		return nil, errors.New("获取用户失败")
	}
	isOk, err := util.ValidatePassword(param.Password, user.HashPassword)
	if !isOk {
		return nil, err
	}
	return &User{
		ID:       user.ID,
		NickName: user.NickName,
		UserName: user.UserName,
	}, nil
}

func (s *UserService) Register(param *UserRegisterRequest) (string, error) {
	hashedPassword, err := util.GeneratePassword(param.Password)
	if err != nil {
		return "", err
	}
	userId, err := s.userDao.Insert(&dao.User{
		ID:           idGenerator.GenerateID(),
		UserName:     param.UserName,
		NickName:     param.NickName,
		HashPassword: string(hashedPassword),
	})
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *UserService) UpdateInfo(param *UserUpdateInfoRequest) error {
	hashedPassword, err := util.GeneratePassword(param.Password)
	if err != nil {
		return err
	}
	err = s.userDao.Update(&dao.User{
		ID:           param.ID,
		UserName:     param.UserName,
		NickName:     param.NickName,
		HashPassword: string(hashedPassword),
	})
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userDao dao.IUser) IUserService {
	return &UserService{userDao: userDao}
}
