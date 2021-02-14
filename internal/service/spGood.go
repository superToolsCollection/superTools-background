package service

import "superTools-background/internal/dao"

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

type ISpGoodService interface {
}

type SpGoodService struct {
	goodDao dao.ISpGood
}

func NewSpGoodService(goodDao dao.ISpGood) ISpGoodService {
	return &SpGoodService{goodDao: goodDao}
}
