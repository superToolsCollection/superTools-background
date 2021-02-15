package service

import (
	"superTools-background/internal/dao"
	"superTools-background/pkg/app"
)

/**
* @Author: super
* @Date: 2021-02-14 20:36
* @Description:
**/
type GetGoodsListRequest struct {
	Query string `form:"query" json:"query"`
}

type AddGoodRequest struct {
	GoodsName      string  `json:"goods_name" form:"cat_name" binding:"required,min=2,max=4294967295"`
	GoodsCat       string  `json:"goods_cat" form:"cat_name" binding:"required,min=2,max=4294967295"`
	GoodsPrice     float64 `json:"goods_price" form:"cat_name" binding:"required,min=2,max=4294967295"`
	GoodsNumber    int     `json:"goods_number" form:"cat_name" binding:"required,min=2,max=4294967295"`
	GoodsWeight    int     `json:"goods_weight" form:"cat_name" binding:"required,min=2,max=4294967295"`
	GoodsIntroduce string  `json:"goods_introduce" form:"goods_introduce"`
	Pics           string  `json:"pics" form:"pics"`
	Attrs          string  `json:"attrs" form:"attrs"`
}

type GetGoodByIdRequest struct {
	ID string `json:"id" form:"id"`
}

type UpdateGoodByIdRequest struct {
	ID string `json:"id" form:"id"`
}

type DeleteGoodByIdRequest struct {
	ID string `json:"id" form:"id"`
}

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

type ISpGoodService interface {
	GetGoodList(param *GetGoodsListRequest, pager *app.Pager) ([]*SpGood, int, error)
}

type SpGoodService struct {
	goodDao dao.ISpGood
}

func (s *SpGoodService) GetGoodList(param *GetGoodsListRequest, pager *app.Pager) ([]*SpGood, int, error) {
	list, totalPage, err := s.goodDao.SelectList(param.Query, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*SpGood, len(list))
	for i, v := range list {
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
	return result, totalPage, nil
}

func NewSpGoodService(goodDao dao.ISpGood) ISpGoodService {
	return &SpGoodService{goodDao: goodDao}
}
