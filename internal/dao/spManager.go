package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
	"superTools-background/pkg/app"
	"time"
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
	SelectList(query string, page int, pageSize int) ([]*SpManager, int, error)
	Insert(spManager *SpManager) (*SpManager, error)
	Update(manager *SpManager) (*SpManager, error)
	UpdateInfo(manager *SpManager) (*SpManager, error)
	Delete(id int) error
}

type SpManagerManger struct {
	table string
	conn  *gorm.DB
}

func (m *SpManagerManger) SelectByID(id int) (*SpManager, error) {
	spManager := &model.SpManager{}
	result := m.conn.Where("mg_id=?", id).Find(spManager)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return &SpManager{
		MgMobile: spManager.MgMobile,
		MgEmail:  spManager.MgEmail,
		MgID:     spManager.MgID,
		MgName:   spManager.MgName,
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

func (m *SpManagerManger) SelectList(query string, page int, pageSize int) ([]*SpManager, int, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	var spManagers []*model.SpManager
	var count int
	stmt := m.conn.Offset(pageOffset).Limit(pageSize).Find(&spManagers)
	if stmt.Error != nil {
		return nil, 0, stmt.Error
	}
	stmt = stmt.Count(&count)
	if stmt.Error != nil {
		return nil, 0, stmt.Error
	}

	result := make([]*SpManager, len(spManagers))
	for i, v := range spManagers {
		result[i] = &SpManager{}
		result[i].MgPwd = v.MgPwd
		result[i].MgEmail = v.MgEmail
		result[i].MgMobile = v.MgMobile
		result[i].MgName = v.MgName
		result[i].RoleID = v.RoleID
		result[i].MgID = v.MgID
		result[i].MgTime = v.MgTime
		result[i].MgState = v.MgState
	}
	return result, count, nil
}

func (m *SpManagerManger) Insert(spManager *SpManager) (*SpManager, error) {
	s := &model.SpManager{
		MgName:   spManager.MgName,
		MgPwd:    spManager.MgPwd,
		MgEmail:  spManager.MgEmail,
		MgMobile: spManager.MgMobile,
		MgTime:   time.Now().Unix(),
		MgState:  1,
	}
	result := m.conn.Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpManager{
		MgID:     s.MgID,
		RoleID:   s.RoleID,
		MgName:   s.MgName,
		MgMobile: s.MgMobile,
		MgEmail:  s.MgEmail,
		MgTime:   s.MgTime,
	}, nil
}

func (m *SpManagerManger) Update(manager *SpManager) (*SpManager, error) {
	s := &model.SpManager{
		MgState: manager.MgState,
	}
	result := m.conn.Model(s).
		Where("mg_id=?", manager.MgID).
		Update("mg_state", manager.MgState)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpManager{
		MgID:     s.MgID,
		RoleID:   s.RoleID,
		MgName:   s.MgName,
		MgMobile: s.MgMobile,
		MgEmail:  s.MgEmail,
		MgTime:   s.MgTime,
	}, nil
}

func (m *SpManagerManger) UpdateInfo(manager *SpManager) (*SpManager, error) {
	s := &model.SpManager{
		MgMobile: manager.MgMobile,
		MgEmail:  manager.MgEmail,
	}
	result := m.conn.Model(s).
		Where("mg_id=?", manager.MgID).Updates(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpManager{
		MgID:     s.MgID,
		RoleID:   s.RoleID,
		MgMobile: s.MgMobile,
		MgEmail:  s.MgEmail,
	}, nil
}

func (m *SpManagerManger) Delete(id int) error {
	result := m.conn.Where("mg_id = ?", id).Delete(model.SpManager{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewSpManagerManger(table string, conn *gorm.DB) ISpManager {
	return &SpManagerManger{table: table, conn: conn}
}
