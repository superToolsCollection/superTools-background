package service

import (
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
	Pid      int    `json:"pid"`
	Level    int    `json:"level"`
	Path     string `json:"path"`
	Children []*SpPermission `json:"children"`
}

type ISpPermissionService interface {
	GetRights(param *GetRightsRequest) ([]*SpPermission, error)
}

type SpPermissionService struct {
	permissionDao    dao.ISpPermission
	permissionApiDao dao.ISpPermissionApi
}

func (s *SpPermissionService) GetRights(param *GetRightsRequest) ([]*SpPermission, error) {
	perList, err := s.permissionDao.Select()
	if err != nil {
		return nil, err
	}
	perApiList, err := s.permissionApiDao.Select()
	if err != nil {
		return nil, err
	}
	perApiMap := make(map[int]string)
	for i := 0; i < len(perApiList); i++ {
		perApiMap[perApiList[i].PsID] = perApiList[i].PsAPIPath
	}
	result := make([]*SpPermission, 0)
	if param.Type == "list" {
		for i := 0; i < len(perList); i++ {
			v := perList[i]
			if v.PsLevel == 0 {
				temp := &SpPermission{
					ID:       v.PsID,
					AuthName: v.PsName,
					Level:    v.PsLevel,
					Pid:      v.PsPid,
					Path:  perApiMap[v.PsID],
				}
				result = append(result, temp)
			}
		}
	} else {
		result = buildPermissionTree(perList, perApiMap)
	}
	return result, nil
}

func buildPermissionTree(perList []*dao.SpPermission, perApiMap map[int]string) []*SpPermission {
	result := make([]*SpPermission, 0)
	for i := 0; i < len(perList); i++ {
		v := perList[i]
		if v.PsLevel == 0 {
			temp := &SpPermission{
				ID:       v.PsID,
				AuthName: v.PsName,
				Level:    v.PsLevel,
				Pid:      v.PsPid,
				Path:     perApiMap[v.PsPid],
				Children: make([]*SpPermission, 0),
			}
			result = append(result, temp)
		}
	}
	level2 := make(map[int]*SpPermission)
	for i := 0; i < len(perList); i++ {
		v := perList[i]
		if v.PsLevel == 1 {
			temp := &SpPermission{
				ID:       v.PsID,
				AuthName: v.PsName,
				Level:    v.PsLevel,
				Pid:      v.PsPid,
				Path:     perApiMap[v.PsPid],
				Children: make([]*SpPermission, 0),
			}
			for j := 0; j < len(result); j++ {
				if result[j].ID == temp.Pid {
					result[j].Children = append(result[j].Children, temp)
					level2[temp.ID] = result[j]
				}
			}
		}
	}
	for i := 0; i < len(perList); i++ {
		v := perList[i]
		if v.PsLevel == 2 {
			temp := &SpPermission{
				ID:       v.PsID,
				AuthName: v.PsName,
				Level:    v.PsLevel,
				Pid:      v.PsPid,
				Path:     perApiMap[v.PsPid],
				Children: nil,
			}
			if v, ok := level2[temp.Pid]; ok {
				for j := 0; j < len(v.Children); j++ {
					if v.Children[j].ID == temp.Pid {
						v.Children[j].Children = append(v.Children[j].Children, temp)
					}
				}
			}
		}
	}
	return result
}

func NewSpPermissionService(permissionDao dao.ISpPermission, permissionApiDao dao.ISpPermissionApi) ISpPermissionService {
	return &SpPermissionService{
		permissionDao:    permissionDao,
		permissionApiDao: permissionApiDao,
	}
}
