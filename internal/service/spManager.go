package service

import (
	"fmt"
	"superTools-background/internal/dao"
	"superTools-background/pkg/app"
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

type GetSpMangerListRequest struct {
	Query string `form:"query"`
}

type AddSpMangerRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=4294967295"`
	Password string `form:"password" binding:"required,min=2,max=4294967295"`
	Email    string `form:"email"`
	Mobile   string `form:"mobile"`
}

type UpdateSpMangerStateRequest struct {
	ID   int64  `form:"id" binding:"required,gte=1"`
	Type string `form:"type" binding:"required"`
}

type GetSpMangerByIDRequest struct {
	ID   int  `form:"id" binding:"required,gte=1"`
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
	GetSpManagerList(param *GetSpMangerListRequest, pager *app.Pager) ([]*SpManager, int, error)
	AddSpManager(param *AddSpMangerRequest) (*SpManager, error)
	UpdateSpManagerState(param *UpdateSpMangerStateRequest) (*SpManager, error)
	GetSpManagerByID(param *GetSpMangerByIDRequest) (*SpManager, error)
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
		MgTime:   spManager.MgTime,
		RoleID:   spManager.RoleID,
	}, nil
}

func (s *SpManagerService) GetSpManagerList(param *GetSpMangerListRequest, pager *app.Pager) ([]*SpManager, int, error) {
	list, totalPage, err := s.dao.SelectList(param.Query, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*SpManager, len(list))
	for i, v := range list {
		result[i] = &SpManager{}
		result[i].MgEmail = v.MgEmail
		result[i].MgMobile = v.MgMobile
		result[i].MgName = v.MgName
		result[i].RoleID = v.RoleID
		result[i].MgID = v.MgID
		result[i].MgTime = v.MgTime
		result[i].MgState = v.MgState
	}
	return result, totalPage, nil
}

func (s *SpManagerService) AddSpManager(param *AddSpMangerRequest) (*SpManager, error) {
	manager := &dao.SpManager{
		MgName:   param.UserName,
		MgPwd:    param.Password,
		MgEmail:  param.Email,
		MgMobile: param.Mobile,
	}
	result, err := s.dao.Insert(manager)
	if err != nil {
		return nil, err
	}
	return &SpManager{
		MgID:     result.MgID,
		RoleID:   result.RoleID,
		MgName:   result.MgName,
		MgMobile: result.MgMobile,
		MgEmail:  result.MgEmail,
		MgTime:   result.MgTime,
	}, nil
}

func (s *SpManagerService) UpdateSpManagerState(param *UpdateSpMangerStateRequest) (*SpManager, error) {
	state := 0
	fmt.Println(param)
	if param.Type == "true" {
		state = 1
	}
	manager := &dao.SpManager{
		MgID:    int(param.ID),
		MgState: state,
	}
	result, err := s.dao.Update(manager)
	if err != nil {
		return nil, err
	}
	return &SpManager{
		MgID:     result.MgID,
		RoleID:   result.RoleID,
		MgName:   result.MgName,
		MgMobile: result.MgMobile,
		MgEmail:  result.MgEmail,
		MgState:  result.MgState,
	}, nil
}

func (s *SpManagerService) GetSpManagerByID(param *GetSpMangerByIDRequest) (*SpManager, error) {
	result, err := s.dao.SelectByID(param.ID)
	if err != nil {
		return nil, err
	}
	return &SpManager{
		MgID:     result.MgID,
		RoleID:   result.RoleID,
		MgName:   result.MgName,
		MgMobile: result.MgMobile,
		MgEmail:  result.MgEmail,
	}, nil
}

func NewSpManagerService(dao dao.ISpManager) ISpManagerService {
	return &SpManagerService{dao: dao}
}
