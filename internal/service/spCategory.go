package service

import (
	"superTools-background/internal/dao"
	"superTools-background/pkg/app"
)

/**
* @Author: super
* @Date: 2021-02-08 14:54
* @Description:
**/
type GetCategoriesListRequest struct {
	Type int `form:"type"`
}

type AddCategoryRequest struct {
	CatPid   int    `form:"cat_pid" binding:"required,gte=1"`
	CatName  string `form:"cat_name" binding:"required,min=2,max=4294967295"`
	CatLevel int    `form:"cat_level" binding:"required,gte=1"`
}

type GetCategoryByIdRequest struct {
	Id int `form:"id"`
}

type SpCategory struct {
	CatID      int           `json:"cat_id"`
	CatName    string        `json:"cat_name"`
	CatPid     int           `json:"cat_pid"`
	CatLevel   int           `json:"cat_level"`
	CatDeleted int           `json:"cat_deleted"`
	Children   []*SpCategory `json:"children"`
}

type ISpCategoryService interface {
	GetCategoriesList(param *GetCategoriesListRequest, pager *app.Pager) ([]*SpCategory, error)
	AddCategory(param *AddCategoryRequest) (*SpCategory, error)
	GetCategory(param *GetCategoryByIdRequest) (*SpCategory, error)
}

type SpCategoryService struct {
	dao dao.ISpCategory
}

func (s *SpCategoryService) GetCategory(param *GetCategoryByIdRequest) (*SpCategory, error) {
	category, err := s.dao.GetCategoryById(param.Id)
	if err != nil {
		return nil, err
	}
	return &SpCategory{
		CatID:    category.CatID,
		CatName:  category.CatName,
		CatPid:   category.CatPid,
		CatLevel: category.CatLevel,
	}, nil
}

func (s *SpCategoryService) AddCategory(param *AddCategoryRequest) (*SpCategory, error) {
	category, err := s.dao.AddCategory(param.CatPid, param.CatName, param.CatLevel)
	if err != nil {
		return nil, err
	}
	return &SpCategory{
		CatID:    category.CatID,
		CatName:  category.CatName,
		CatPid:   category.CatPid,
		CatLevel: category.CatLevel,
	}, nil
}

func (s *SpCategoryService) GetCategoriesList(param *GetCategoriesListRequest, pager *app.Pager) ([]*SpCategory, error) {
	categories, err := s.dao.SelectCategoriesList(param.Type, pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	result := buildTree(categories, param.Type)
	return result, nil
}

func buildTree(list []*dao.SpCategory, treeType int) []*SpCategory {
	result := make([]*SpCategory, 0)
	for i := 0; i < len(list); i++ {
		v := list[i]
		if v.CatLevel == 0 {
			temp := &SpCategory{
				CatID:      v.CatID,
				CatName:    v.CatName,
				CatPid:     v.CatPid,
				CatLevel:   v.CatLevel,
				CatDeleted: v.CatDeleted,
				Children:   make([]*SpCategory, 0),
			}
			result = append(result, temp)
		}
	}
	if treeType == 1 {
		return result
	}
	level2 := make(map[int]*SpCategory)
	for i := 0; i < len(list); i++ {
		v := list[i]
		if v.CatLevel == 1 {
			temp := &SpCategory{
				CatID:      v.CatID,
				CatName:    v.CatName,
				CatPid:     v.CatPid,
				CatLevel:   v.CatLevel,
				CatDeleted: v.CatDeleted,
				Children:   make([]*SpCategory, 0),
			}
			for j := 0; j < len(result); j++ {
				if result[j].CatID == temp.CatPid {
					result[j].Children = append(result[j].Children, temp)
					level2[temp.CatID] = result[j]
				}
			}
		}
	}
	if treeType == 2 {
		return result
	}
	for i := 0; i < len(list); i++ {
		v := list[i]
		if v.CatLevel == 2 {
			temp := &SpCategory{
				CatID:      v.CatID,
				CatName:    v.CatName,
				CatPid:     v.CatPid,
				CatLevel:   v.CatLevel,
				CatDeleted: v.CatDeleted,
				Children:   nil,
			}
			if v, ok := level2[temp.CatPid]; ok {
				for j := 0; j < len(v.Children); j++ {
					if v.Children[j].CatID == temp.CatPid {
						v.Children[j].Children = append(v.Children[j].Children, temp)
					}
				}
			}
		}
	}
	return result
}

func NewSpCategoryService(dao dao.ISpCategory) ISpCategoryService {
	return &SpCategoryService{dao: dao}
}
