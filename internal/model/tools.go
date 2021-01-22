package model

/**
* @Author: super
* @Date: 2020-11-25 13:28
* @Description:
**/

type Tool struct {
	*Model
	APIDescribe string `gorm:"column:api_describe" json:"api_describe"`
	Name        string `gorm:"column:name" json:"name"`
	State       int    `gorm:"column:state;force" json:"state"`
	API         string `gorm:"column:api" json:"api"`
}

// TableName sets the insert table name for this struct type
func (t *Tool) TableName() string {
	return "tools"
}
