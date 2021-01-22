package model

/**
* @Author: super
* @Date: 2021-01-22 20:17
* @Description: 属性表
**/

type SpAttribute struct {
	AttrID     int    `gorm:"column:attr_id;primary_key" json:"attr_id"`
	AttrName   string `gorm:"column:attr_name" json:"attr_name"`
	CatID      int    `gorm:"column:cat_id" json:"cat_id"`
	AttrSel    string `gorm:"column:attr_sel" json:"attr_sel"`
	AttrWrite  string `gorm:"column:attr_write" json:"attr_write"`
	AttrVals   string `gorm:"column:attr_vals" json:"attr_vals"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"`
}

// TableName sets the insert table name for this struct type
func (s *SpAttribute) TableName() string {
	return "sp_attribute"
}
