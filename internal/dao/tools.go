package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
	"superTools-background/pkg/app"
)

/**
* @Author: super
* @Date: 2020-11-25 13:35
* @Description:
**/

type Tool struct {
	CreateOn    string `json:"create_on"`
	CreatedBy   string `json:"created_by"`
	ID          string `json:"id"`
	ModifiedBy  string `json:"modified_by"`
	APIDescribe string `json:"api_describe"`
	DeleteOn    string `json:"delete_on"`
	ModifiedOn  string `json:"modified_on"`
	Name        string `json:"name"`
	State       int    `json:"state"`
	API         string `json:"api"`
}

type ITool interface {
	Insert(tool *Tool) (string, error)
	Delete(id string) bool
	Update(tool *Tool) error
	SelectByKey(id string) (*model.Tool, error)
	SelectAll() ([]*model.Tool, error)
	SelectByName(name string) (*model.Tool, error)
	SelectList(page, pageSize int) ([]*model.Tool, error)
}

type ToolManager struct {
	table string
	conn  *gorm.DB
}

func (m *ToolManager) Insert(tool *Tool) (string, error) {
	t := &model.Tool{
		Model: &model.Model{
			CreatedBy: tool.CreatedBy,
		},
		ID:          tool.ID,
		APIDescribe: tool.APIDescribe,
		Name:        tool.Name,
		State:       tool.State,
		API:         tool.API,
	}
	result := m.conn.Create(t)
	if result.RowsAffected == int64(0) {
		return "", errors.New("insert error")
	}
	return t.ID, nil
}

func (m *ToolManager) Delete(id string) bool {
	result := m.conn.Where("id=?", id).Delete(model.Tool{})
	if result.RowsAffected == int64(0) {
		return false
	}
	return true
}

func (m *ToolManager) Update(tool *Tool) error {
	t := &model.Tool{
		Model: &model.Model{
			ModifiedBy: tool.ModifiedBy,
		},
		ID:          tool.ID,
		APIDescribe: tool.APIDescribe,
		Name:        tool.Name,
		State:       tool.State,
		API:         tool.API,
	}
	result := m.conn.Model(t).Where("id=?", t.ID).Updates(t)
	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}
	return nil
}

func (m *ToolManager) SelectByKey(id string) (*model.Tool, error) {
	t := &model.Tool{}
	result := m.conn.Where("id=?", id).Find(t)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return t, nil
}

func (m *ToolManager) SelectAll() ([]*model.Tool, error) {
	var tools []*model.Tool
	if err := m.conn.Find(&tools).Error; err != nil {
		return nil, err
	}
	return tools, nil
}

func (m *ToolManager) SelectByName(name string) (*model.Tool, error) {
	t := &model.Tool{}
	result := m.conn.Where("name=?", name).Find(t)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return t, nil
}

func (m *ToolManager) SelectList(page, pageSize int) ([]*model.Tool, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	fields := []string{"id", "name", "api", "api_describe", "created_on", "created_by", "modified_on", "modified_by", "state"}
	rows, err := m.conn.Offset(pageOffset).Limit(pageSize).Select(fields).Table(m.table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tools []*model.Tool
	for rows.Next() {
		tool := &model.Tool{}
		if err := rows.Scan(&tool.ID,
			&tool.Name,
			&tool.API,
			&tool.APIDescribe,
			&tool.CreatedOn,
			&tool.CreatedBy,
			&tool.ModifiedOn,
			&tool.ModifiedBy,
			&tool.State); err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}
	return tools, nil
}

func NewToolManager(table string, conn *gorm.DB) ITool {
	return &ToolManager{table: table, conn: conn}
}
