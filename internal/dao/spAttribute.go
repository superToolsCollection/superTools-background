package dao

import (
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
)

/**
* @Author: super
* @Date: 2021-02-13 19:48
* @Description:
**/

type SpAttribute struct {
	AttrID     int    `json:"attr_id"`
	AttrName   string `json:"attr_name"`
	CatID      int    `json:"cat_id"`
	AttrSel    string `json:"attr_sel"`
	AttrWrite  string `json:"attr_write"`
	AttrVals   string `json:"attr_vals"`
	DeleteTime int64  `json:"delete_time"`
}

type ISpAttribute interface {
	GetAttribute(catID int, sel string) (*SpAttribute, error)
	AddAttribute(catId int, name string, sel string, vals string) (*SpAttribute, error)
	DeleteAttribute(catId int, attrId int) error
	GetAttributeById(catId int, attrId int) (*SpAttribute, error)
	UpdateAttribute(catId int, attrId int, name string, sel string, vals string) (*SpAttribute, error)
}

type SpAttributeManager struct {
	table string
	conn  *gorm.DB
}

func (m *SpAttributeManager) UpdateAttribute(catId int, attrId int, name string, sel string, vals string) (*SpAttribute, error) {
	attribute := &model.SpAttribute{}
	result := m.conn.Model(attribute).
		Where("cat_id=? and attr_id=?", catId, attrId).
		Update("cat_name=?", name).
		Update("cat_sel=?", sel).
		Update("cat_vals=?", vals)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpAttribute{
		CatID:     attribute.CatID,
		AttrID:    attribute.AttrID,
		AttrName:  attribute.AttrName,
		AttrSel:   attribute.AttrSel,
		AttrWrite: attribute.AttrWrite,
		AttrVals:  attribute.AttrVals,
	}, nil
}

func (m *SpAttributeManager) GetAttributeById(catId int, attrId int) (*SpAttribute, error) {
	attribute := &model.SpAttribute{}
	result := m.conn.Where("cat_id=? and attr_id=?", catId, attrId).Find(attribute)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpAttribute{
		CatID:     attribute.CatID,
		AttrID:    attribute.AttrID,
		AttrName:  attribute.AttrName,
		AttrSel:   attribute.AttrSel,
		AttrWrite: attribute.AttrWrite,
		AttrVals:  attribute.AttrVals,
	}, nil
}

func (m *SpAttributeManager) DeleteAttribute(catId int, attrId int) error {
	result := m.conn.Where("cat_id=? and attr_id=?", catId, attrId).Delete(&model.SpAttribute{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *SpAttributeManager) AddAttribute(catId int, name string, sel string, vals string) (*SpAttribute, error) {
	attribute := &model.SpAttribute{
		CatID:    catId,
		AttrName: name,
		AttrSel:  sel,
		AttrVals: vals,
	}
	result := m.conn.Create(attribute)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpAttribute{
		CatID:     attribute.CatID,
		AttrID:    attribute.AttrID,
		AttrName:  attribute.AttrName,
		AttrSel:   attribute.AttrSel,
		AttrWrite: attribute.AttrWrite,
		AttrVals:  attribute.AttrVals,
	}, nil
}

func (m *SpAttributeManager) GetAttribute(catID int, sel string) (*SpAttribute, error) {
	attribute := &model.SpAttribute{}
	result := m.conn.Where("cat_id=?", catID).Find(attribute)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SpAttribute{
		CatID:     attribute.CatID,
		AttrID:    attribute.AttrID,
		AttrName:  attribute.AttrName,
		AttrSel:   attribute.AttrSel,
		AttrWrite: attribute.AttrWrite,
		AttrVals:  attribute.AttrVals,
	}, nil
}

func NewSpAttributeManager(table string, conn *gorm.DB) ISpAttribute {
	return &SpAttributeManager{table: table, conn: conn}
}
