package dao

import (
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2021-02-07 20:09
* @Description:
**/

type SpRole struct {
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	PsIds    string `json:"ps_ids"`
	PsCa     string `json:"ps_ca"`
	RoleDesc string `json:"role_desc"`
}

type ISpRole interface {
	Select() ([]*SpRole, error)
}

type SpRoleManager struct {
	table string
	conn  *gorm.DB
}

func NewSpRoleManager(table string, conn *gorm.DB) ISpRole {
	return &SpRoleManager{table: table, conn: conn}
}

func (m *SpRoleManager) Select() ([]*SpRole, error) {
	var roleList []*model.SpRole
	stmt := m.conn.Find(&roleList)
	if stmt.Error != nil {
		return nil, stmt.Error
	}
	result := make([]*SpRole, len(roleList))
	for i := 0; i < len(roleList); i++ {
		result[i] = &SpRole{}
		result[i].RoleID = roleList[i].RoleID
		result[i].PsCa = roleList[i].PsCa
		result[i].PsIds = roleList[i].PsIds
		result[i].RoleDesc = roleList[i].RoleDesc
		result[i].RoleName = roleList[i].RoleName
	}
	return result, nil
}
