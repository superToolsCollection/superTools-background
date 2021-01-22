package model

/**
* @Author: super
* @Date: 2021-01-22 20:24
* @Description: 快递表
**/

type SpExpress struct {
	ExpressNu  string `gorm:"column:express_nu" json:"express_nu"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	ExpressID  int    `gorm:"column:express_id;primary_key" json:"express_id"`
	OrderID    int    `gorm:"column:order_id" json:"order_id"`
	ExpressCom string `gorm:"column:express_com" json:"express_com"`
}

// TableName sets the insert table name for this struct type
func (s *SpExpress) TableName() string {
	return "sp_express"
}
