package service

import (
	"superTools-background/internal/dao"
	"superTools-background/pkg/app"
	"superTools-background/pkg/idGenerator"
)

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
	Name        string `form:"name" binding:"required,min=2,max=4294967295"`
	ModifiedBy  string `form:"modified_by" binding:"required,min=2,max=4294967295"`
	API         string `form:"api" binding:"required,min=2,max=4294967295"`
	APIDescribe string `form:"api_describe" binding:"required,min=2,max=4294967295"`
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
	ToolOnLine(param *ToolOnLineRequest) error
	ToolOffLine(param *ToolOffLineRequest) error
	GetToolByKey(param *GetToolByKeyRequest) (*Tool, error)
	GetToolByName(param *GetToolByNameRequest) (*Tool, error)
	GetToolList(pager *app.Pager) ([]*Tool, error)
}

type ToolService struct {
	toolDao dao.ITool
}

func NewToolService(toolDao dao.ITool) IToolService {
	return &ToolService{toolDao: toolDao}
}

func (s *ToolService) AddTool(param *AddToolRequest) (string, error) {
	tool := &dao.Tool{
		ID:          idGenerator.GenerateID(),
		Name:        param.Name,
		APIDescribe: param.APIDescribe,
		API:         param.API,
		State:       0,
		CreatedBy:   param.CreatedBy,
	}
	return s.toolDao.Insert(tool)
}

func (s *ToolService) UpdateToolInfo(param *UpdateToolRequest) error {
	tool := &dao.Tool{
		ID:          param.ID,
		Name:        param.Name,
		APIDescribe: param.APIDescribe,
		API:         param.API,
		ModifiedBy:  param.ModifiedBy,
	}
	return s.toolDao.Update(tool)
}

func (s *ToolService) DeleteTool(param *DeleteToolRequest) bool {
	return s.toolDao.Delete(param.ID)
}

func (s *ToolService) ToolOnLine(param *ToolOnLineRequest) error {
	tool := &dao.Tool{
		ID:    param.ID,
		State: 1,
	}
	err := s.toolDao.Update(tool)
	if err != nil {
		return err
	}
	return nil
}

func (s *ToolService) ToolOffLine(param *ToolOffLineRequest) error {
	tool := &dao.Tool{
		ID:    param.ID,
		State: 0,
	}
	err := s.toolDao.Update(tool)
	if err != nil {
		return err
	}
	return nil
}

func (s *ToolService) GetToolByKey(param *GetToolByKeyRequest) (*Tool, error) {
	result, err := s.toolDao.SelectByKey(param.ID)
	if err != nil {
		return nil, err
	}
	tool := &Tool{
		ID:          result.ID,
		Name:        result.Name,
		API:         result.API,
		APIDescribe: result.APIDescribe,
		CreatedBy:   result.CreatedBy,
		CreateOn:    result.CreatedOn,
		ModifiedBy:  result.ModifiedBy,
		ModifiedOn:  result.ModifiedOn,
		DeleteOn:    result.DeletedOn,
		State:       result.State,
	}
	return tool, nil
}

func (s *ToolService) GetToolByName(param *GetToolByNameRequest) (*Tool, error) {
	result, err := s.toolDao.SelectByName(param.Name)
	if err != nil {
		return nil, err
	}
	tool := &Tool{
		ID:          result.ID,
		Name:        result.Name,
		API:         result.API,
		APIDescribe: result.APIDescribe,
		CreatedBy:   result.CreatedBy,
		CreateOn:    result.CreatedOn,
		ModifiedBy:  result.ModifiedBy,
		ModifiedOn:  result.ModifiedOn,
		DeleteOn:    result.DeletedOn,
		State:       result.State,
	}
	return tool, nil
}

func (s *ToolService) GetToolList(pager *app.Pager) ([]*Tool, error) {
	tools, err := s.toolDao.SelectList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	var toolList []*Tool
	for _, tool := range tools {
		toolList = append(toolList, &Tool{
			ID:          tool.ID,
			Name:        tool.Name,
			API:         tool.API,
			APIDescribe: tool.APIDescribe,
			CreatedBy:   tool.CreatedBy,
			CreateOn:    tool.CreatedOn,
			ModifiedBy:  tool.ModifiedBy,
			ModifiedOn:  tool.ModifiedOn,
			DeleteOn:    tool.DeletedOn,
			State:       tool.State,
		})
	}
	return toolList, nil
}
