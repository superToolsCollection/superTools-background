package mall

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/convert"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-11-21 15:57
* @Description:
**/

type OrderController struct {
	OrderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) OrderController {
	return OrderController{OrderService: orderService}
}

// @Summary 获取订单列表
// @Tags mall
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [get]
func (o OrderController) GetOrderList(c *gin.Context) {
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	orders, err := o.OrderService.GetOrderList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListFail)
		return
	}
	data := gin.H{
		"page_num":  pager.Page,
		"page_size": pager.PageSize,
		"orders":    orders,
	}
	response.ToResponse(data, "获取订单列表成功", http.StatusOK)
	return
}

// @Summary 获取所有订单
// @Tags mall
// @Produce json
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/all_orders [get]
func (o OrderController) GetAllOrder(c *gin.Context) {
	response := app.NewResponse(c)
	orders, err := o.OrderService.GetAllOrder()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetAllOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAllOrderFail)
		return
	}
	response.ToResponse(orders, "获取全部清单成功", http.StatusOK)
	return
}

// @Summary 获取用户所有订单
// @Tags mall
// @Produce json
// @Param user_id query int false "用户id"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/all_orders_user [get]
func (o OrderController) GetOrderListByUserID(c *gin.Context) {
	param := service.GetOrderListByUserIDRequest{}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	orders, err := o.OrderService.GetOrderListByUserID(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderListByUserID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListByUserIDFail)
		return
	}
	data := gin.H{
		"page_num":  pager.Page,
		"page_size": pager.PageSize,
		"orders":    orders,
	}
	response.ToResponse(data, "获取用户订单列表成功", http.StatusOK)
	return
}

// @Summary 获取用户订单列表
// @Tags mall
// @Produce json
// @Param user_id query int false "用户id"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders_user [get]
func (o OrderController) GetOrderByUserID(c *gin.Context) {
	param := service.GetOrderByUserIDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	orders, err := o.OrderService.GetOrderByUserID(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderByUserID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderByUserIDFail)
		return
	}
	response.ToResponse(orders, "获取用户订单成功", http.StatusOK)
	return
}

// @Summary 获取单个订单
// @Tags mall
// @Produce json
// @Param id path int true "订单ID"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders/{id} [get]
func (o OrderController) GetOrder(c *gin.Context) {
	param := service.OrderRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	order, err := o.OrderService.GetOrderByID(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderByID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderFail)
		return
	}
	response.ToResponse(order, "获取订单成功", http.StatusOK)
	return
}

// @Summary 新增订单
// @Tags mall
// @Produce json
// @Param id body int true "订单id"
// @Param user_id body int true "用户id"
// @Param product_id body int true "商品id"
// @Param state body int true "订单状态"
// @Success 200 {object} int "1"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [post]
func (o OrderController) Insert(c *gin.Context) {
	param := service.InsertOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	_, err := o.OrderService.InsertOrder(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.InsertOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorInsertOrderFail)
		return
	}
	response.ToResponse(gin.H{}, "插入订单成功", http.StatusCreated)
	return
}

// @Summary 更新订单
// @Tags mall
// @Produce json
// @Param id body int true "订单id"
// @Param user_id body int true "用户id"
// @Param product_id body int true "商品id"
// @Param state body int true "订单状态"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [put]
func (o OrderController) Update(c *gin.Context) {
	param := service.UpdateOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := o.OrderService.UpdateOrder(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateOrderFail)
		return
	}
	response.ToResponse(gin.H{}, "更新订单成功", http.StatusOK)
	return
}

// @Summary 删除订单
// @Tags mall
// @Produce json
// @Param id path int true "订单ID"
// @Success 200 {string} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders/{id} [delete]
func (o OrderController) Delete(c *gin.Context) {
	param := service.DeleteOrderRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	result := o.OrderService.DeleteOrderByID(&param)
	if result != true {
		global.Logger.Errorf(c, "svc.DeleteOrderByID err: %v", result)
		response.ToErrorResponse(errcode.ErrorDeleteOrderFail)
		return
	}
	response.ToResponse(gin.H{}, "删除订单成功", http.StatusOK)
	return
}
