package model

/**
* @Author: super
* @Date: 2021-01-22 20:21
* @Description: 收货人信息
**/

type SpConsignee struct {
	CgnAddress string `gorm:"column:cgn_address" json:"cgn_address"`
	CgnTel     string `gorm:"column:cgn_tel" json:"cgn_tel"`
	CgnCode    string `gorm:"column:cgn_code" json:"cgn_code"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"`
	CgnID      int    `gorm:"column:cgn_id;primary_key" json:"cgn_id"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
	CgnName    string `gorm:"column:cgn_name" json:"cgn_name"`
}

// TableName sets the insert table name for this struct type
func (s *SpConsignee) TableName() string {
	return "sp_consignee"
}
