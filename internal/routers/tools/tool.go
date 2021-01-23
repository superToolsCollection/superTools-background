package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-11-25 22:31
* @Description:
**/

type ToolController struct {
	ToolService service.IToolService
}

func NewToolController(toolService service.IToolService) ToolController {
	return ToolController{ToolService: toolService}
}

// @Summary 新增工具
// @Tags tool
// @Produce json
// @Param name body string true "工具名称"
// @Param created_by body string true "创建者"
// @Param api body string true "工具api地址"
// @Param api_describe body string true "工具api描述"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/addTool [post]
func (t ToolController) AddTool(c *gin.Context) {
	param := service.AddToolRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	_, err := t.ToolService.AddTool(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.AddTool err: %v", err)
		response.ToErrorResponse(errcode.ErrorAddToolFail)
		return
	}
	response.ToResponse(gin.H{}, "添加工具成功", http.StatusOK)
	return
}

// @Summary 更新工具信息
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Param name body string true "工具名称"
// @Param modified_by body string true "修改者"
// @Param api body string true "工具api地址"
// @Param api_describe body string true "工具api描述"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/update [put]
func (t ToolController) UpdateToolInfo(c *gin.Context) {
	param := service.UpdateToolRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := t.ToolService.UpdateToolInfo(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateToolInfo err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateToolInfoFail)
		return
	}
	response.ToResponse(gin.H{}, "更新工具信息成功", http.StatusOK)
	return
}

// @Summary 删除工具
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/delete [post]
func (t ToolController) DeleteTool(c *gin.Context) {
	param := service.DeleteToolRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := t.ToolService.DeleteTool(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTool err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteToolFail)
		return
	}
	response.ToResponse(gin.H{}, "删除工具成功", http.StatusOK)
	return
}

// @Summary 上线工具
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/toolOnLine [put]
func (t ToolController) ToolOnLine(c *gin.Context) {
	param := service.ToolOnLineRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := t.ToolService.ToolOnLine(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.ToolOnLine err: %v", err)
		response.ToErrorResponse(errcode.ErrorToolOnlineFail)
		return
	}
	response.ToResponse(gin.H{}, "工具上线成功", http.StatusOK)
	return
}

// @Summary 下线工具
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/toolOffLine [put]
func (t ToolController) ToolOffLine(c *gin.Context) {
	param := service.ToolOffLineRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := t.ToolService.ToolOffLine(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.ToolOffLine err: %v", err)
		response.ToErrorResponse(errcode.ErrorToolOffLineFail)
		return
	}
	response.ToResponse(gin.H{}, "工具下线成功", http.StatusOK)
	return
}

// @Summary 根据id获取工具
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/getTool [get]
func (t ToolController) GetToolByKey(c *gin.Context) {
	param := service.GetToolByKeyRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tool, err := t.ToolService.GetToolByKey(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolByKey err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolByKeyFail)
		return
	}
	response.ToResponse(tool, "获取工具成功", http.StatusOK)
	return
}

// @Summary 根据名称获取工具
// @Tags tool
// @Produce json
// @Param name body string true "工具名称"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/getToolByName [get]
func (t ToolController) GetToolByName(c *gin.Context) {
	param := service.GetToolByNameRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tool, err := t.ToolService.GetToolByName(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolByName err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolByNameFail)
		return
	}
	response.ToResponse(tool, "获取工具成功", http.StatusOK)
	return
}

// @Summary 获取工具列表
// @Tags tool
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/toolList [get]
func (t ToolController) GetToolList(c *gin.Context) {
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	tools, err := t.ToolService.GetToolList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolListFail)
		return
	}
	data := gin.H{
		"page_num":pager.Page,
		"page_size":pager.PageSize,
		"tools":tools,
	}
	response.ToResponse(data, "获取工具列表成功", http.StatusOK)
	return
}
