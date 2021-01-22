package model

/**
* @Author: super
* @Date: 2021-01-22 20:32
* @Description: 商品订单关联表
**/

type SpOrderGood struct {
	GoodsID         int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsPrice      float64 `gorm:"column:goods_price" json:"goods_price"`
	GoodsNumber     int     `gorm:"column:goods_number" json:"goods_number"`
	GoodsTotalPrice float64 `gorm:"column:goods_total_price" json:"goods_total_price"`
	ID              int     `gorm:"column:id;primary_key" json:"id"`
	OrderID         int     `gorm:"column:order_id" json:"order_id"`
}

// TableName sets the insert table name for this struct type
func (s *SpOrderGood) TableName() string {
	return "sp_order_goods"
}
