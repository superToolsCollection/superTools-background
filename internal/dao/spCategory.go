package dao

import (
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
	"superTools-background/pkg/app"
)

/**
* @Author: super
* @Date: 2021-02-08 14:50
* @Description:
**/

type SpCategory struct {
	CatID      int    `json:"cat_id"`
	CatName    string `json:"cat_name"`
	CatPid     int    `json:"cat_pid"`
	CatLevel   int    `json:"cat_level"`
	CatDeleted int    `json:"cat_deleted"`
	CatIcon    string `json:"cat_icon"`
	CatSrc     string `json:"cat_src"`
}

type ISpCategory interface {
	SelectCategoriesList(treeType int, page int, pageSize int) ([]*SpCategory, error)
	AddCategory(pid int, name string, level int) (*SpCategory, error)
	GetCategoryById(id int) (*SpCategory, error)
	UpdateCategory(id int, name string) (*SpCategory, error)
	DeleteCategory(id int) error
}

type SpCategoryManager struct {
	table string
	conn  *gorm.DB
}

func (m *SpCategoryManager) DeleteCategory(id int) error {
	result := m.conn.Where("cat_id = ?", id).Delete(model.SpCategory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *SpCategoryManager) UpdateCategory(id int, name string) (*SpCategory, error) {
	c := &model.SpCategory{}
	result := m.conn.Model(c).Where("cat_id=?", id).Update("cat_name", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpCategory{
		CatID:    c.CatID,
		CatName:  c.CatName,
		CatPid:   c.CatPid,
		CatLevel: c.CatLevel,
	}, nil
}

func (m *SpCategoryManager) GetCategoryById(id int) (*SpCategory, error) {
	c := &model.SpCategory{}
	result := m.conn.Where("cat_id=?", id).Find(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpCategory{
		CatID:    c.CatID,
		CatName:  c.CatName,
		CatPid:   c.CatPid,
		CatLevel: c.CatLevel,
	}, nil
}

func (m *SpCategoryManager) AddCategory(pid int, name string, level int) (*SpCategory, error) {
	c := &model.SpCategory{
		CatPid:   pid,
		CatName:  name,
		CatLevel: level,
	}
	result := m.conn.Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpCategory{
		CatID:    c.CatID,
		CatName:  c.CatName,
		CatPid:   c.CatPid,
		CatLevel: c.CatLevel,
	}, nil
}

func (m *SpCategoryManager) SelectCategoriesList(treeType int, page int, pageSize int) ([]*SpCategory, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	var daoCategories []*model.SpCategory
	if pageOffset <= 0 && pageSize <= 0 {
		stmt := m.conn.Find(&daoCategories)
		if stmt.Error != nil {
			return nil, stmt.Error
		}
	} else {
		stmt := m.conn.Offset(pageOffset).Limit(pageSize).Find(&daoCategories)
		if stmt.Error != nil {
			return nil, stmt.Error
		}
	}
	result := make([]*SpCategory, len(daoCategories))
	for i := 0; i < len(daoCategories); i++ {
		result[i] = &SpCategory{
			CatPid:     daoCategories[i].CatPid,
			CatID:      daoCategories[i].CatID,
			CatLevel:   daoCategories[i].CatLevel,
			CatDeleted: daoCategories[i].CatDeleted,
			CatName:    daoCategories[i].CatName,
			CatSrc:     daoCategories[i].CatSrc,
			CatIcon:    daoCategories[i].CatIcon,
		}
	}
	return result, nil
}

func NewSpCategoryManager(table string, conn *gorm.DB) ISpCategory {
	return &SpCategoryManager{table: table, conn: conn}
}
