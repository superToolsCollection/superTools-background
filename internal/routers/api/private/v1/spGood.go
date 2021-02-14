package v1

import "superTools-background/internal/service"

/**
* @Author: super
* @Date: 2021-02-14 20:48
* @Description:
**/

type SpGoodController struct {
	GoodService service.ISpGoodService
}

func NewSpGoodController(goodService service.ISpGoodService) SpGoodController {
	return SpGoodController{GoodService: goodService}
}
