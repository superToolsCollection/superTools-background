package model

/**
* @Author: super
* @Date: 2021-01-22 20:42
* @Description: 用户购物车表
**/

type SpUserCart struct {
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"`
	CartID     int    `gorm:"column:cart_id;primary_key" json:"cart_id"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
	CartInfo   string `gorm:"column:cart_info" json:"cart_info"`
}

// TableName sets the insert table name for this struct type
func (s *SpUserCart) TableName() string {
	return "sp_user_cart"
}
