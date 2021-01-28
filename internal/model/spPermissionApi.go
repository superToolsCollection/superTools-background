package model

/**
* @Author: super
* @Date: 2021-01-22 20:35
* @Description: api描述表
**/

type SpPermissionApi struct {
	PsAPIAction  string `gorm:"column:ps_api_action" json:"ps_api_action"`
	PsAPIPath    string `gorm:"column:ps_api_path" json:"ps_api_path"`
	PsAPIOrder   int    `gorm:"column:ps_api_order" json:"ps_api_order"`
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	PsID         int    `gorm:"column:ps_id" json:"ps_id"`
	PsAPIService string `gorm:"column:ps_api_service" json:"ps_api_service"`
}

// TableName sets the insert table name for this struct type
func (s *SpPermissionApi) TableName() string {
	return "sp_permission_api"
}
