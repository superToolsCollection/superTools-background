package dao

import "github.com/jinzhu/gorm"

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

}

type SpCategoryManager struct {
	table string
	conn *gorm.DB
}

func NewSpCategoryManager(table string, conn *gorm.DB) ISpCategory {
	return &SpCategoryManager{table:table, conn:conn}
}