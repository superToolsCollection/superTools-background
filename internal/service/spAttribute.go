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
	Id     int `form:"id"`
	AttrId int `form:"attr_id"`
}

type ISpAttributeService interface {
}

type SpAttributeService struct {
	dao dao.ISpAttribute
}

func NewSpAttribute(dao dao.ISpAttribute) ISpAttributeService {
	return &SpAttributeService{dao: dao}
}
