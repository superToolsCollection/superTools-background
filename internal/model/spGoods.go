package model

/**
* @Author: super
* @Date: 2021-01-22 20:25
* @Description: 商品表
**/

type SpGood struct {
	DeleteTime     int64   `gorm:"column:delete_time" json:"delete_time"`
	CatThreeID     int     `gorm:"column:cat_three_id" json:"cat_three_id"`
	CatTwoID       int     `gorm:"column:cat_two_id" json:"cat_two_id"`
	// 0: 未通过 1: 审核中 2: 已审核
	GoodsState     int     `gorm:"column:goods_state" json:"goods_state"`
	GoodsID        int     `gorm:"column:goods_id;primary_key" json:"goods_id"`
	GoodsWeight    int     `gorm:"column:goods_weight" json:"goods_weight"`
	GoodsBigLogo   string  `gorm:"column:goods_big_logo" json:"goods_big_logo"`
	UpdTime        int64   `gorm:"column:upd_time" json:"upd_time"`
	CatOneID       int     `gorm:"column:cat_one_id" json:"cat_one_id"`
	HotMumber      int     `gorm:"column:hot_mumber" json:"hot_mumber"`
	AddTime        int64   `gorm:"column:add_time" json:"add_time"`
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsPrice     float64 `gorm:"column:goods_price" json:"goods_price"`
	GoodsNumber    int     `gorm:"column:goods_number" json:"goods_number"`
	CatID          int     `gorm:"column:cat_id" json:"cat_id"`
	GoodsIntroduce string  `gorm:"column:goods_introduce" json:"goods_introduce"`
	GoodsSmallLogo string  `gorm:"column:goods_small_logo" json:"goods_small_logo"`
	IsDel          int  `gorm:"column:is_del" json:"is_del"`
	IsPromote      int     `gorm:"column:is_promote" json:"is_promote"`
}

// TableName sets the insert table name for this struct type
func (s *SpGood) TableName() string {
	return "sp_goods"
}
