package service

import (
	"strconv"
	"strings"
	"superTools-background/internal/dao"
)

/**
* @Author: super
* @Date: 2021-02-07 20:24
* @Description:
**/
type AddRoleRequest struct {
	RoleName string `form:"role_name" binding:"required,min=2,max=4294967295"`
	RoleDesc string `form:"role_desc"`
}

type GetRoleByIdRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}

type UpdateRoleRequest struct {
	ID       int    `form:"id"`
	RoleName string `form:"role_name" binding:"required,min=2,max=4294967295"`
	RoleDesc string `form:"role_desc"`
}

type DeleteRoleRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}

type UpdateRightRequest struct {
	RoleID int    `form:"roleId"`
	Rids   string `form:"rids" binding:"required,min=2,max=4294967295"`
}

type DeleteRightRequest struct {
	RoleID  int `form:"roleId"`
	RightId int `form:"rightId"`
}

type SpRole struct {
	RoleID   int             `json:"id"`
	RoleName string          `json:"roleName"`
	Children []*SpPermission `json:"children"`
	PsCa     string          `json:"ps_ca"`
	RoleDesc string          `json:"role_desc"`
}

type ISpRoleService interface {
	GetRoles() ([]*SpRole, error)
	GetRoleByID(param *GetRoleByIdRequest) (*SpRole, error)
	AddRole(param *AddRoleRequest) (*SpRole, error)
	UpdateRole(param *UpdateRoleRequest) (*SpRole, error)
	DeleteRole(param *DeleteRoleRequest) error
	UpdateRight(param *UpdateRightRequest) error
	DeleteRight(param *DeleteRightRequest) error
}

type SpRoleService struct {
	roleDao       dao.ISpRole
	permissionDao dao.ISpPermission
}

func NewSpRoleService(roleDao dao.ISpRole, permissionDao dao.ISpPermission) ISpRoleService {
	return &SpRoleService{
		roleDao:       roleDao,
		permissionDao: permissionDao,
	}
}

func (s *SpRoleService) GetRoles() ([]*SpRole, error) {
	roleList, err := s.roleDao.Select()
	if err != nil {
		return nil, err
	}
	result := make([]*SpRole, len(roleList))
	for i := 0; i < len(roleList); i++ {
		idStrs := strings.Split(roleList[i].PsIds, ",")
		ids := make([]int, len(idStrs))
		for j := 0; j < len(idStrs); j++ {
			t, err := strconv.Atoi(idStrs[j])
			if err != nil {
				continue
			}
			ids[j] = t
		}
		perList, err := s.permissionDao.SelectByIds(ids)
		if err != nil {
			return nil, err
		}
		//todo 会有bug，map不能为空
		perTree := buildPermissionTree(perList, nil)
		temp := &SpRole{
			RoleName: roleList[i].RoleName,
			RoleID:   roleList[i].RoleID,
			Children: perTree,
			PsCa:     roleList[i].PsCa,
			RoleDesc: roleList[i].RoleDesc,
		}
		result[i] = temp
	}
	return result, nil
}

func (s *SpRoleService) GetRoleByID(param *GetRoleByIdRequest) (*SpRole, error) {
	result, err := s.roleDao.SelectByID(param.ID)
	if err != nil {
		return nil, err
	}
	return &SpRole{
		RoleID:   result.RoleID,
		RoleName: result.RoleName,
		RoleDesc: result.RoleDesc,
	}, nil
}

func (s *SpRoleService) AddRole(param *AddRoleRequest) (*SpRole, error) {
	role := &dao.SpRole{
		RoleName: param.RoleName,
		RoleDesc: param.RoleDesc,
	}
	result, err := s.roleDao.Insert(role)
	if err != nil {
		return nil, err
	}
	return &SpRole{
		RoleDesc: result.RoleDesc,
		RoleID:   result.RoleID,
		RoleName: result.RoleName,
	}, nil
}

func (s *SpRoleService) UpdateRole(param *UpdateRoleRequest) (*SpRole, error) {
	role := &dao.SpRole{
		RoleID:   param.ID,
		RoleName: param.RoleName,
		RoleDesc: param.RoleDesc,
	}
	result, err := s.roleDao.Update(role)
	if err != nil {
		return nil, err
	}
	return &SpRole{
		RoleDesc: result.RoleDesc,
		RoleID:   result.RoleID,
		RoleName: result.RoleName,
	}, nil
}

func (s *SpRoleService) DeleteRole(param *DeleteRoleRequest) error {
	role := &dao.SpRole{
		RoleID: param.ID,
	}
	err := s.roleDao.Delete(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpRoleService) UpdateRight(param *UpdateRightRequest) error {
	role := &dao.SpRole{
		RoleID: param.RoleID,
		PsIds:  param.Rids,
	}
	err := s.roleDao.UpdateRight(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpRoleService) DeleteRight(param *DeleteRightRequest) error {
	role, err := s.roleDao.SelectByID(param.RoleID)
	if err != nil {
		return err
	}
	idStrs := strings.Split(role.PsIds, ",")
	ids := make([]string, 0)
	for j := 0; j < len(idStrs); j++ {
		t := idStrs[j]
		if t == strconv.Itoa(param.RightId) {
			continue
		}
		ids = append(ids, t)
	}
	psids := strings.Join(ids, ",")
	role.PsIds = psids
	err = s.roleDao.UpdateRight(role)
	if err != nil {
		return err
	}
	return nil
}
