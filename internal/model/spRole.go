package model

/**
* @Author: super
* @Date: 2021-01-22 20:39
* @Description: 角色表
**/

type SpRole struct {
	RoleID   int    `gorm:"column:role_id;primary_key" json:"role_id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
	PsIds    string `gorm:"column:ps_ids" json:"ps_ids"`
	PsCa     string `gorm:"column:ps_ca" json:"ps_ca"`
	RoleDesc string `gorm:"column:role_desc" json:"role_desc"`
}

// TableName sets the insert table name for this struct type
func (s *SpRole) TableName() string {
	return "sp_role"
}
