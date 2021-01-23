package bedtimeStory

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
* @Date: 2020-09-16 07:51
* @Description: story对应的restful api
**/

type StoryController struct {
	StoryService service.IStoryService
}

func NewStoryController(storyService service.IStoryService) StoryController {
	return StoryController{StoryService: storyService}
}

// @Summary 随机获取单个故事
// @tags 睡前故事
// @Produce json
// @Success 200 {object} model.Story "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/story [get]
func (t StoryController) Get(c *gin.Context) {
	response := app.NewResponse(c)
	story, err := t.StoryService.GetStory()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoryFail)
		return
	}

	response.ToResponse(story, "获取故事成功", http.StatusOK)
	return
}
