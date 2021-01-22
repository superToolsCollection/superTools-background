package model

/**
* @Author: super
* @Date: 2021-01-22 20:33
* @Description: 权限表
**/

type SpPermission struct {
	PsID    int    `gorm:"column:ps_id;primary_key" json:"ps_id"`
	PsName  string `gorm:"column:ps_name" json:"ps_name"`
	PsPid   int    `gorm:"column:ps_pid" json:"ps_pid"`
	PsC     string `gorm:"column:ps_c" json:"ps_c"`
	PsA     string `gorm:"column:ps_a" json:"ps_a"`
	PsLevel int    `gorm:"column:ps_level" json:"ps_level"`
}

// TableName sets the insert table name for this struct type
func (s *SpPermission) TableName() string {
	return "sp_permission"
}
