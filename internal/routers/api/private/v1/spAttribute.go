package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/convert"
	"superTools-background/pkg/errcode"
)

/**
* @Author: super
* @Date: 2021-02-13 19:47
* @Description:
**/

type SpAttributeController struct {
	Service service.ISpAttributeService
}

func NewSpAttributeController(service service.ISpAttributeService) SpAttributeController {
	return SpAttributeController{Service: service}
}

func (s SpAttributeController) GetAttribute(c *gin.Context) {
	param := service.GetAttributeRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.GetAttribute errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorGetCategoryFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	attribute, err := s.Service.GetAttribute(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpAttributeService.GetAttribute errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAttributeFail)
		return
	}
	response.ToResponse(attribute, "获取成功", http.StatusOK)
}

func (s SpAttributeController) AddAttribute(c *gin.Context) {
	param := service.AddAttributeRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.AddAttribute errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorAddAttributeFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	attribute, err := s.Service.AddAttribute(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpAttributeService.AddAttribute errs: %v", err)
		response.ToErrorResponse(errcode.ErrorAddAttributeFail)
		return
	}
	response.ToResponse(attribute, "创建成功", http.StatusCreated)
}

func (s SpAttributeController) DeleteAttribute(c *gin.Context) {
	param := service.DeleteAttributeRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.DeleteAttribute errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorDeleteAttributeFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	attrIdStr := strings.TrimSpace(c.Param("attrid"))
	if attrIdStr == "" || len(attrIdStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.DeleteAttribute errs: %v", errors.New("wrong attr_id"))
		response.ToErrorResponse(errcode.ErrorDeleteAttributeFail)
		return
	}
	attrId := convert.StrTo(attrIdStr).MustInt64()
	param.AttrId = int(attrId)

	err := s.Service.DeleteAttribute(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpAttributeService.DeleteAttribute errs: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteAttributeFail)
		return
	}
	response.ToResponse(gin.H{}, "删除成功", http.StatusCreated)
}

func (s SpAttributeController) GetAttributeById(c *gin.Context) {
	param := service.GetAttributeByIdRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.GetAttributeById errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorGetAttributeByIdFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	attrIdStr := strings.TrimSpace(c.Param("attrid"))
	if attrIdStr == "" || len(attrIdStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.GetAttributeById errs: %v", errors.New("wrong attr_id"))
		response.ToErrorResponse(errcode.ErrorGetAttributeByIdFail)
		return
	}
	attrId := convert.StrTo(attrIdStr).MustInt64()
	param.AttrId = int(attrId)

	attribute, err := s.Service.GetAttributeById(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpAttributeService.GetAttributeById errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAttributeByIdFail)
		return
	}
	response.ToResponse(attribute, "获取成功", http.StatusOK)
}

func (s SpAttributeController) UpdateAttribute(c *gin.Context) {
	param := service.UpdateAttributeRequest{}
	response := app.NewResponse(c)
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" || len(idStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.UpdateAttribute errs: %v", errors.New("wrong id"))
		response.ToErrorResponse(errcode.ErrorUpdateAttributeByIdFail)
		return
	}
	id := convert.StrTo(idStr).MustInt64()
	param.Id = int(id)
	attrIdStr := strings.TrimSpace(c.Param("attrid"))
	if attrIdStr == "" || len(attrIdStr) == 0 {
		global.Logger.Errorf(c, "SpAttributeService.UpdateAttribute errs: %v", errors.New("wrong attr_id"))
		response.ToErrorResponse(errcode.ErrorUpdateAttributeByIdFail)
		return
	}
	attrId := convert.StrTo(attrIdStr).MustInt64()
	param.AttrId = int(attrId)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	attribute, err := s.Service.UpdateAttribute(&param)
	if err != nil {
		global.Logger.Errorf(c, "SpAttributeService.UpdateAttribute errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAttributeByIdFail)
		return
	}
	response.ToResponse(attribute, "更新成功", http.StatusOK)
}
