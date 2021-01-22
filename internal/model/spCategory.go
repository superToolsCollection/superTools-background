package model

/**
* @Author: super
* @Date: 2021-01-22 20:21
* @Description: 分类表
**/

type SpCategory struct {
	CatID      int    `gorm:"column:cat_id;primary_key" json:"cat_id"`
	CatName    string `gorm:"column:cat_name" json:"cat_name"`
	CatPid     int    `gorm:"column:cat_pid" json:"cat_pid"`
	CatLevel   int    `gorm:"column:cat_level" json:"cat_level"`
	CatDeleted int    `gorm:"column:cat_deleted" json:"cat_deleted"`
	CatIcon    string `gorm:"column:cat_icon" json:"cat_icon"`
	CatSrc     string `gorm:"column:cat_src" json:"cat_src"`
}

// TableName sets the insert table name for this struct type
func (s *SpCategory) TableName() string {
	return "sp_category"
}
