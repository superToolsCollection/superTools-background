package model

/**
* @Author: super
* @Date: 2021-01-22 20:28
* @Description: 商品-相册关联表
**/

type SpGoodsPics struct {
	PicsID  int    `gorm:"column:pics_id;primary_key" json:"pics_id"`
	GoodsID int    `gorm:"column:goods_id" json:"goods_id"`
	PicsBig string `gorm:"column:pics_big" json:"pics_big"`
	PicsMid string `gorm:"column:pics_mid" json:"pics_mid"`
	PicsSma string `gorm:"column:pics_sma" json:"pics_sma"`
}

// TableName sets the insert table name for this struct type
func (s *SpGoodsPics) TableName() string {
	return "sp_goods_pics"
}
