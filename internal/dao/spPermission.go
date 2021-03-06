package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2021-01-28 17:20
* @Description:
**/

type SpPermission struct {
	PsID    int    `json:"ps_id"`
	PsName  string `json:"ps_name"`
	PsPid   int    `json:"ps_pid"`
	PsC     string `json:"ps_c"`
	PsA     string `json:"ps_a"`
	PsLevel int    `json:"ps_level"`
}

func (s SpPermission) String() string {
	return fmt.Sprintf("ps_id: %d, "+
		"ps_name:%s, "+
		"ps_pid:%d, "+
		"ps_c:%s, "+
		"ps_a:%s, "+
		"ps_level:%d\n",
		s.PsID, s.PsName, s.PsPid, s.PsC, s.PsA, s.PsLevel)
}

type ISpPermission interface {
	Select() ([]*SpPermission, error)
	SelectByIds(ids []int) ([]*SpPermission, error)
}

type SpPermissionManger struct {
	table string
	conn  *gorm.DB
}

func (m *SpPermissionManger) Select() ([]*SpPermission, error) {
	var permissionList []*model.SpPermission
	stmt := m.conn.Find(&permissionList)
	if stmt.Error != nil {
		return nil, stmt.Error
	}
	result := make([]*SpPermission, len(permissionList))
	for i := 0; i < len(permissionList); i++ {
		result[i] = &SpPermission{}
		result[i].PsID = permissionList[i].PsID
		result[i].PsName = permissionList[i].PsName
		result[i].PsLevel = permissionList[i].PsLevel
		result[i].PsPid = permissionList[i].PsPid
	}
	return result, nil
}

func (m *SpPermissionManger) SelectByIds(ids []int) ([]*SpPermission, error) {
	var permissionList []*model.SpPermission
	stmt := m.conn.Where("ps_id IN (?)", ids).Find(&permissionList)
	if stmt.Error != nil {
		return nil, stmt.Error
	}
	result := make([]*SpPermission, len(permissionList))
	for i := 0; i < len(permissionList); i++ {
		result[i] = &SpPermission{}
		result[i].PsID = permissionList[i].PsID
		result[i].PsName = permissionList[i].PsName
		result[i].PsLevel = permissionList[i].PsLevel
		result[i].PsPid = permissionList[i].PsPid
	}
	return result, nil
}

func NewSpPermissionManger(table string, conn *gorm.DB) ISpPermission {
	return &SpPermissionManger{table: table, conn: conn}
}
