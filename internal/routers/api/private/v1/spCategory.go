package v1

import (
	"github.com/gin-gonic/gin"
	"superTools-background/internal/service"
)

/**
* @Author: super
* @Date: 2021-02-08 14:55
* @Description:
**/

type SpCategoryController struct {
	Service service.ISpCategoryService
}

func NewSpCategoryController(service service.ISpCategoryService) SpCategoryController{
	return SpCategoryController{Service:service}
}

func (s SpCategoryController)GetCateforiesList(c *gin.Context){

}