package model

/**
* @Author: super
* @Date: 2021-01-22 20:40
* @Description: 类型表
**/

type SpType struct {
	TypeID     int    `gorm:"column:type_id;primary_key" json:"type_id"`
	TypeName   string `gorm:"column:type_name" json:"type_name"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"`
}

// TableName sets the insert table name for this struct type
func (s *SpType) TableName() string {
	return "sp_type"
}
