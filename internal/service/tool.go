package service

import "superTools-background/internal/dao"

/**
* @Author: super
* @Date: 2020-11-25 15:06
* @Description:
**/
type AddToolRequest struct {
	Name        string `form:"name" binding:"required,min=2,max=4294967295"`
	CreatedBy   string `form:"created_by" binding:"required,min=2,max=4294967295"`
	API         string `form:"api" binding:"required,min=2,max=4294967295"`
	APIDescribe string `form:"api_describe" binding:"required,min=2,max=4294967295"`
}

type UpdateToolRequest struct {
	ID          string `form:"id" binding:"required,min=2,max=4294967295"`
	Name        string `form:"user_id" binding:"required,min=2,max=4294967295"`
	ModifiedBy  string `form:"user_id" binding:"required,min=2,max=4294967295"`
	API         string `form:"product_id" binding:"required,min=2,max=4294967295"`
	APIDescribe string `form:"product_id" binding:"required,min=2,max=4294967295"`
}

type DeleteToolRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

type ToolOnLineRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

type ToolOffLineRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

type GetToolByKeyRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

type GetToolByNameRequest struct {
	Name string `form:"name" binding:"required,min=2,max=4294967295"`
}

type GetToolListRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

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

type IToolService interface {
	AddTool(param *AddToolRequest) (string, error)
	UpdateToolInfo(param *UpdateToolRequest) error
	DeleteTool(param *DeleteToolRequest) bool
	ToolOnLine(param *ToolOnLineRequest) bool
	ToolOffLine(param *ToolOffLineRequest) bool
	GetToolByKey(param *GetToolByKeyRequest) (*Tool, error)
	GetToolByName(param *GetToolByNameRequest) (*Tool, error)
	GetToolList(param *GetToolListRequest) ([]*Tool, error)
}

type ToolService struct {
	toolDao dao.ITool
}

func NewToolService(toolDao dao.ITool) IToolService {
	return &ToolService{toolDao: toolDao}
}

func (*ToolService) AddTool(param *AddToolRequest) (string, error) {
	panic("implement me")
}

func (*ToolService) UpdateToolInfo(param *UpdateToolRequest) error {
	panic("implement me")
}

func (*ToolService) DeleteTool(param *DeleteToolRequest) bool {
	panic("implement me")
}

func (*ToolService) ToolOnLine(param *ToolOnLineRequest) bool {
	panic("implement me")
}

func (*ToolService) ToolOffLine(param *ToolOffLineRequest) bool {
	panic("implement me")
}

func (*ToolService) GetToolByKey(param *GetToolByKeyRequest) (*Tool, error) {
	panic("implement me")
}

func (*ToolService) GetToolByName(param *GetToolByNameRequest) (*Tool, error) {
	panic("implement me")
}

func (*ToolService) GetToolList(param *GetToolListRequest) ([]*Tool, error) {
	panic("implement me")
}
