package service

import (
	"superTools-background/internal/dao"
)

/**
* @Author: super
* @Date: 2021-02-08 14:54
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

type ISpCategoryService interface {

}

type SpCategoryService struct {
	dao dao.ISpCategory
}

func NewSpCategoryService(dao dao.ISpCategory) ISpCategoryService {
	return &SpCategoryService{dao:dao}
}