package service

import (
	"fmt"
	"superTools-background/internal/dao"
)

/**
* @Author: super
* @Date: 2021-01-28 17:20
* @Description:
**/

type GetRightsRequest struct {
	Type string `form:"user_name" binding:"required,min=2,max=4294967295"`
}

type SpPermission struct {
	ID       int    `json:"id"`
	AuthName string `json:"authName"`
	Pid      string `json:"pid"`
	Level    int    `json:"level"`
	Path     string `json:"path"`
	Children []SpPermission
}

type ISpPermissionService interface {
	GetRights(param *GetRightsRequest) ([]SpPermission, error)
}

type SpPermissionService struct {
	permissionDao    dao.ISpPermission
	permissionApiDao dao.ISpPermissionApi
}

func (s *SpPermissionService) GetRights(param *GetRightsRequest) ([]SpPermission, error) {
	perList, err := s.permissionDao.Select()
	if err != nil {
		return nil, err
	}
	perMap := make(map[int]*dao.SpPermission)
	for i := 0; i < len(perList); i++ {
		perMap[perList[i].PsID] = perList[i]
	}
	perApiList, err := s.permissionApiDao.Select()
	if err != nil {
		return nil, err
	}
	perApiMap := make(map[int]*dao.SpPermissionApi)
	for i := 0; i < len(perApiList); i++ {
		perApiMap[perApiList[i].PsID] = perApiList[i]
	}
	result := make([]SpPermission, 0)
	if param.Type == "list" {
		for i := 0; i < len(perList); i++ {
			v := perList[i]
			if v.PsLevel == 0 {
				temp := SpPermission{
					ID:       v.PsID,
					AuthName: v.PsName,
					Level:    v.PsLevel,
					Pid:      fmt.Sprintf("%d", v.PsPid),
				}
				result = append(result, temp)
			}
		}
	} else {
		//todo 完善tree权限列表
		for i := 0; i < len(perList); i++ {
			v := perList[i]
			if v.PsLevel == 0 {
				temp := SpPermission{
					ID:       v.PsID,
					AuthName: v.PsName,
					Level:    v.PsLevel,
					Pid:      fmt.Sprintf("%d", v.PsPid),
					Path:     perApiMap[v.PsPid].PsAPIPath,
					Children: make([]SpPermission, 0),
				}
				result = append(result, temp)
			}
		}
	}
	return result, nil
}

func NewSpPermissionService(permissionDao dao.ISpPermission, permissionApiDao dao.ISpPermissionApi) ISpPermissionService {
	return &SpPermissionService{
		permissionDao:    permissionDao,
		permissionApiDao: permissionApiDao,
	}
}
