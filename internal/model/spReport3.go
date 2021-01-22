package model

/**
* @Author: super
* @Date: 2021-01-22 20:38
* @Description:
**/

type SpReport3 struct {
	ID       int    `gorm:"column:id;primary_key" json:"id"`
	Rp3Src   string `gorm:"column:rp3_src" json:"rp3_src"`
	Rp3Count int    `gorm:"column:rp3_count" json:"rp3_count"`
	// 数据库中是date类型
	Rp3Date string `gorm:"column:rp3_date" json:"rp3_date"`
}

// TableName sets the insert table name for this struct type
func (s *SpReport3) TableName() string {
	return "sp_report_3"
}
