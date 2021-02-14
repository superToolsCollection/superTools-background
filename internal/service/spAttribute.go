package service

import "superTools-background/internal/dao"

/**
* @Author: super
* @Date: 2021-02-13 19:50
* @Description:
**/
type GetAttributeRequest struct {
	Id  int    `form:"id"`
	Sel string `form:"sel" binding:"required,oneof=only many"`
}

type AddAttributeRequest struct {
	Id       int    `form:"id"`
	AttrName string `form:"attr_name" binding:"required,min=2,max=4294967295"`
	AttrSel  string `form:"attr_sel" binding:"required,oneof=only many"`
	AttrVals string `form:"attr_vals"`
}

type DeleteAttributeRequest struct {
	Id     int `form:"id"`
	AttrId int `form:"attr_id"`
}

type GetAttributeByIdRequest struct {
	Id     int `form:"id"`
	AttrId int `form:"attr_id"`
}

type UpdateAttributeRequest struct {
	Id       int    `form:"id"`
	AttrId   int    `form:"attr_id"`
	AttrName string `form:"attr_name" binding:"required,min=2,max=4294967295"`
	AttrSel  string `form:"attr_sel" binding:"required,oneof=only many"`
	AttrVals string `form:"attr_vals"`
}

type SpAttribute struct {
	AttrID     int    `json:"attr_id"`
	AttrName   string `json:"attr_name"`
	CatID      int    `json:"cat_id"`
	AttrSel    string `json:"attr_sel"`
	AttrWrite  string `json:"attr_write"`
	AttrVals   string `json:"attr_vals"`
	DeleteTime int64  `json:"delete_time"`
}

type ISpAttributeService interface {
	GetAttribute(param *GetAttributeRequest) (*SpAttribute, error)
	AddAttribute(param *AddAttributeRequest) (*SpAttribute, error)
	DeleteAttribute(param *DeleteAttributeRequest) error
	GetAttributeById(param *GetAttributeByIdRequest) (*SpAttribute, error)
	UpdateAttribute(param *UpdateAttributeRequest) (*SpAttribute, error)
}

type SpAttributeService struct {
	dao dao.ISpAttribute
}

func (s *SpAttributeService) UpdateAttribute(param *UpdateAttributeRequest) (*SpAttribute, error) {
	attribute, err := s.dao.UpdateAttribute(param.Id, param.AttrId, param.AttrName, param.AttrSel, param.AttrVals)
	if err != nil {
		return nil, err
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

func (s *SpAttributeService) GetAttributeById(param *GetAttributeByIdRequest) (*SpAttribute, error) {
	attribute, err := s.dao.GetAttributeById(param.Id, param.AttrId)
	if err != nil {
		return nil, err
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

func (s *SpAttributeService) DeleteAttribute(param *DeleteAttributeRequest) error {
	err := s.dao.DeleteAttribute(param.Id, param.AttrId)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpAttributeService) AddAttribute(param *AddAttributeRequest) (*SpAttribute, error) {
	attribute, err := s.dao.AddAttribute(param.Id, param.AttrName, param.AttrSel, param.AttrVals)
	if err != nil {
		return nil, err
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

func (s *SpAttributeService) GetAttribute(param *GetAttributeRequest) (*SpAttribute, error) {
	attribute, err := s.dao.GetAttribute(param.Id, param.Sel)
	if err != nil {
		return nil, err
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

func NewSpAttribute(dao dao.ISpAttribute) ISpAttributeService {
	return &SpAttributeService{dao: dao}
}
