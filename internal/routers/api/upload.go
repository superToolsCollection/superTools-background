package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"superTools-background/global"
	"superTools-background/internal/service"
	"superTools-background/pkg/app"
	"superTools-background/pkg/convert"
	"superTools-background/pkg/errcode"
	"superTools-background/pkg/upload"
)

/**
* @Author: super
* @Date: 2020-09-23 19:10
* @Description:
**/

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 用于文件上传
// @Produce json
// @Param file formData file true "文件"
// @Param type formData int true "文件类型"
// @Success 200 {string} string "file_url"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"temp_path": fileInfo.TempPath,
		"url":       fileInfo.Url,
	}, "上传成功", http.StatusOK)
}
