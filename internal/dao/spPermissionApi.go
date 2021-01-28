package dao

import (
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2021-01-28 18:13
* @Description:
**/

type SpPermissionApi struct {
	PsAPIAction  string `json:"ps_api_action"`
	PsAPIPath    string `json:"ps_api_path"`
	PsAPIOrder   int    `json:"ps_api_order"`
	ID           int    `json:"id"`
	PsID         int    `json:"ps_id"`
	PsAPIService string `json:"ps_api_service"`
}

type ISpPermissionApi interface {
	Select() ([]*SpPermissionApi, error)
}

type SpPermissionApiManger struct {
	table string
	conn  *gorm.DB
}

func (m *SpPermissionApiManger) Select() ([]*SpPermissionApi, error) {
	var permissionList []*model.SpPermissionApi
	stmt := m.conn.Find(&permissionList)
	if stmt.Error != nil {
		return nil, stmt.Error
	}
	result := make([]*SpPermissionApi, len(permissionList))
	for i := 0; i < len(permissionList); i++ {
		result[i] = &SpPermissionApi{}
		result[i].PsID = permissionList[i].PsID
		result[i].PsAPIAction = permissionList[i].PsAPIAction
		result[i].PsAPIPath = permissionList[i].PsAPIPath
		result[i].PsAPIOrder = permissionList[i].PsAPIOrder
		result[i].ID = permissionList[i].ID
		result[i].PsAPIService = permissionList[i].PsAPIService
	}
	return result, nil
}

func NewSpPermissionApiManger(table string, conn *gorm.DB) ISpPermissionApi {
	return &SpPermissionApiManger{table: table, conn: conn}
}
