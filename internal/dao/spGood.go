package dao

import (
	"github.com/jinzhu/gorm"
	"superTools-background/internal/model"
	"superTools-background/pkg/app"
)

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
	SelectList(query string, page int, pageSize int) ([]*SpGood, int, error)
}

type SpGoodManager struct {
	table string
	conn  *gorm.DB
}

func (m *SpGoodManager) SelectList(query string, page int, pageSize int) ([]*SpGood, int, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	var spGoods []*model.SpGood
	var count int
	stmt := m.conn.Offset(pageOffset).Limit(pageSize).Find(&spGoods).Count(&count)
	if stmt.Error != nil {
		return nil, 0, stmt.Error
	}
	//stmt = stmt.Count(&count)
	//if stmt.Error != nil {
	//	return nil, 0, stmt.Error
	//}

	result := make([]*SpGood, len(spGoods))
	for i, v := range spGoods {
		result[i] = &SpGood{}
		result[i].GoodsID = v.GoodsID
		result[i].GoodsName = v.GoodsName
		result[i].GoodsPrice = v.GoodsPrice
		result[i].GoodsNumber = v.GoodsNumber
		result[i].GoodsWeight = v.GoodsWeight
		result[i].GoodsState = v.GoodsState
		result[i].AddTime = v.AddTime
		result[i].UpdTime = v.UpdTime
		result[i].HotMumber = v.HotMumber
		result[i].IsPromote = v.IsPromote
	}
	return result, count, nil
}

func NewSpGoodManager(table string, conn *gorm.DB) ISpGood {
	return &SpGoodManager{table: table, conn: conn}
}
