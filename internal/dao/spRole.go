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
	SelectByID(id int) (*SpRole, error)
	Insert(role *SpRole) (*SpRole, error)
	Update(role *SpRole) (*SpRole, error)
	Delete(role *SpRole) error
	UpdateRight(role *SpRole) error
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

func (m *SpRoleManager) SelectByID(id int) (*SpRole, error) {
	var role *model.SpRole
	stmt := m.conn.Where("role_id=?", id).Find(&role)
	if stmt.Error != nil {
		return nil, stmt.Error
	}
	return &SpRole{
		RoleID:   role.RoleID,
		RoleName: role.RoleName,
		PsIds:role.PsIds,
		RoleDesc: role.RoleDesc,
	}, nil
}

func (m *SpRoleManager) Insert(role *SpRole) (*SpRole, error) {
	s := &model.SpRole{
		RoleDesc: role.RoleDesc,
		RoleID:   role.RoleID,
		RoleName: role.RoleName,
		PsIds:    role.PsIds,
		PsCa:     role.PsCa,
	}
	result := m.conn.Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpRole{
		RoleDesc: s.RoleDesc,
		RoleID:   s.RoleID,
		RoleName: s.RoleName,
		PsIds:    s.PsIds,
		PsCa:     s.PsCa,
	}, nil
}

func (m *SpRoleManager) Update(role *SpRole) (*SpRole, error) {
	r := &model.SpRole{}
	result := m.conn.Model(r).
		Where("role_id=?", role.RoleID).
		Update("role_name", role.RoleName).
		Update("role_desc", role.RoleDesc)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpRole{
		RoleID:   r.RoleID,
		RoleName: r.RoleName,
		RoleDesc: r.RoleDesc,
	}, nil
}

func (m *SpRoleManager) Delete(role *SpRole) error {
	result := m.conn.Where("role_id = ?", role.RoleID).Delete(model.SpManager{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *SpRoleManager) UpdateRight(role *SpRole) error{
	r := &model.SpRole{}
	result := m.conn.Model(r).
		Where("role_id=?", role.RoleID).
		Update("ps_ids", role.PsIds)
	if result.Error != nil {
		return result.Error
	}
	return nil
}