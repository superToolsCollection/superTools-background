package model

/**
* @Author: super
* @Date: 2021-01-22 20:36
* @Description:
**/

type SpReport1 struct {
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Rp1UserCount int    `gorm:"column:rp1_user_count" json:"rp1_user_count"`
	Rp1Area      string `gorm:"column:rp1_area" json:"rp1_area"`
	// 数据库中是Date类型
	Rp1Date string `gorm:"column:rp1_date" json:"rp1_date"`
}

// TableName sets the insert table name for this struct type
func (s *SpReport1) TableName() string {
	return "sp_report_1"
}
