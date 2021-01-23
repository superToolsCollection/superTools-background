package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2021-01-23 20:10
* @Description:
**/

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

type ISpManager interface {
	SelectByID(id int) (*SpManager, error)
	SelectByNamePwd(name string) (*SpManager, error)
}

type SpManagerManger struct {
	table string
	conn  *gorm.DB
}

func (m *SpManagerManger) SelectByID(id int) (*SpManager, error) {
	spManager := &model.SpManager{}
	result := m.conn.Where("id=?", id).Find(spManager)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
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

func (m *SpManagerManger) SelectByNamePwd(name string) (*SpManager, error) {
	spManger := &model.SpManager{}
	result := m.conn.Where("mg_name=?", name).Find(spManger)
	if result.RecordNotFound() {
		return nil, errors.New("wrong name or pwd")
	}
	return &SpManager{
		MgMobile: spManger.MgMobile,
		MgEmail:  spManger.MgEmail,
		MgState:  spManger.MgState,
		MgID:     spManger.MgID,
		MgName:   spManger.MgName,
		MgPwd:    spManger.MgPwd,
		MgTime:   spManger.MgTime,
		RoleID:   spManger.RoleID,
	}, nil
}

func NewSpManagerManger(table string, conn *gorm.DB) ISpManager {
	return &SpManagerManger{table: table, conn: conn}
}
