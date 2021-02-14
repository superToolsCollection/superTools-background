package dao

import "github.com/jinzhu/gorm"

/**
* @Author: super
* @Date: 2021-02-14 20:36
* @Description:
**/

type SpGood struct {
	DeleteTime     int64   `json:"delete_time"`
	CatThreeID     int     `json:"cat_three_id"`
	CatTwoID       int     `json:"cat_two_id"`
	GoodsState     int     `json:"goods_state"`
	GoodsID        int     `json:"goods_id"`
	GoodsWeight    int     `json:"goods_weight"`
	GoodsBigLogo   string  `json:"goods_big_logo"`
	UpdTime        int64   `json:"upd_time"`
	CatOneID       int     `json:"cat_one_id"`
	HotMumber      int     `json:"hot_mumber"`
	AddTime        int64   `json:"add_time"`
	GoodsName      string  `json:"goods_name"`
	GoodsPrice     float64 `json:"goods_price"`
	GoodsNumber    int     `json:"goods_number"`
	CatID          int     `json:"cat_id"`
	GoodsIntroduce string  `json:"goods_introduce"`
	GoodsSmallLogo string  `json:"goods_small_logo"`
	IsDel          int     `json:"is_del"`
	IsPromote      int     `json:"is_promote"`
}

type ISpGood interface {
}

type SpGoodManager struct {
	table string
	conn  *gorm.DB
}

func NewSpGoodManager(table string, conn *gorm.DB) ISpGood {
	return &SpGoodManager{table: table, conn: conn}
}
