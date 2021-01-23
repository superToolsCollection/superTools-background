package service

import (
	"superTools-background/internal/dao"
	"superTools-background/pkg/util"
)

/**
* @Author: super
* @Date: 2021-01-23 20:29
* @Description:
**/

type GetSpMangerRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
}

type SpManager struct {
	MgMobile string `json:"mg_mobile"`
	MgEmail  string `json:"mg_email"`
	MgState  int    `json:"mg_state"`
	MgID     int    `json:"mg_id"`
	MgName   string `json:"mg_name"`
	MgPwd    string `json:"mg_pwd"`
	MgTime   int64  `json:"mg_time"`
	RoleID   int    `json:"role_id"`
}

type ISpManagerService interface {
	GetSpManager(param *GetSpMangerRequest) (*SpManager, error)
}

type SpManagerService struct {
	dao dao.ISpManager
}

func (s *SpManagerService) GetSpManager(param *GetSpMangerRequest) (*SpManager, error) {
	spManager, err := s.dao.SelectByNamePwd(param.UserName)
	if err != nil {
		return nil, err
	}
	isOk, err := util.ValidatePassword(param.Password, spManager.MgPwd)
	if !isOk {
		return nil, err
	}
	return &SpManager{
		MgMobile: spManager.MgMobile,
		MgEmail:  spManager.MgEmail,
		MgState:  spManager.MgState,
		MgID:     spManager.MgID,
		MgName:   spManager.MgName,
		MgPwd:    spManager.MgPwd,
		MgTime:   spManager.MgTime,
		RoleID:   spManager.RoleID,
	}, nil
}

func NewSpManagerService(dao dao.ISpManager) ISpManagerService {
	return &SpManagerService{dao: dao}
}
