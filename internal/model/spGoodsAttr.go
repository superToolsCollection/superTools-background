package model

/**
* @Author: super
* @Date: 2021-01-22 20:26
* @Description: 商品-属性关联表
**/

type SpGoodsAttr struct {
	GoodsID   int     `gorm:"column:goods_id" json:"goods_id"`
	AttrID    int     `gorm:"column:attr_id" json:"attr_id"`
	AttrValue string  `gorm:"column:attr_value" json:"attr_value"`
	AddPrice  float64 `gorm:"column:add_price" json:"add_price"`
	ID        int     `gorm:"column:id;primary_key" json:"id"`
}

// TableName sets the insert table name for this struct type
func (s *SpGoodsAttr) TableName() string {
	return "sp_goods_attr"
}
