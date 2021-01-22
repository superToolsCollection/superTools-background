package model

/**
* @Author: super
* @Date: 2021-01-22 20:28
* @Description: 商品分类表
**/

type SpGoodsCat struct {
	CatSort    int    `gorm:"column:cat_sort" json:"cat_sort"`
	DataFlag   int    `gorm:"column:data_flag" json:"data_flag"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	CatID      int    `gorm:"column:cat_id;primary_key" json:"cat_id"`
	ParentID   int    `gorm:"column:parent_id" json:"parent_id"`
	CatName    string `gorm:"column:cat_name" json:"cat_name"`
	IsShow     int    `gorm:"column:is_show" json:"is_show"`
}

// TableName sets the insert table name for this struct type
func (s *SpGoodsCat) TableName() string {
	return "sp_goods_cats"
}
