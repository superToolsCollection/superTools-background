package model

/**
* @Author: super
* @Date: 2021-01-22 20:29
* @Description: 管理员表
**/

type SpManager struct {
	MgMobile string `gorm:"column:mg_mobile" json:"mg_mobile"`
	MgEmail  string `gorm:"column:mg_email" json:"mg_email"`
	MgState  int    `gorm:"column:mg_state" json:"mg_state"`
	MgID     int    `gorm:"column:mg_id;primary_key" json:"mg_id"`
	MgName   string `gorm:"column:mg_name" json:"mg_name"`
	MgPwd    string `gorm:"column:mg_pwd" json:"mg_pwd"`
	MgTime   int64  `gorm:"column:mg_time" json:"mg_time"`
	RoleID   int    `gorm:"column:role_id" json:"role_id"`
}

// TableName sets the insert table name for this struct type
func (s *SpManager) TableName() string {
	return "sp_manager"
}
