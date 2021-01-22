package model

/**
* @Author: super
* @Date: 2021-01-22 20:37
* @Description:
**/

type SpReport2 struct {
	ID       int    `gorm:"column:id;primary_key" json:"id"`
	Rp2Page  string `gorm:"column:rp2_page" json:"rp2_page"`
	Rp2Count int    `gorm:"column:rp2_count" json:"rp2_count"`
	// 数据库中是date类型
	Rp2Date string `gorm:"column:rp2_date" json:"rp2_date"`
}

// TableName sets the insert table name for this struct type
func (s *SpReport2) TableName() string {
	return "sp_report_2"
}
