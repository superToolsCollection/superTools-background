package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "superTools-background/docs"
	"superTools-background/global"
	"superTools-background/internal/dao"
	"superTools-background/internal/middleware"
	"superTools-background/internal/routers/api/private/v1"
	"superTools-background/internal/service"
	"superTools-background/pkg/limiter"
	"time"
)

/**
* @Author: super
* @Date: 2021-01-23 19:11
* @Description:
**/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	spManager := dao.NewSpManagerManger("sp_manager", global.DBEngine)
	spService := service.NewSpManagerService(spManager)
	spController := v1.NewSpManagerController(spService)

	userGroup := r.Group("/api/private/v1/")
	{
		userGroup.POST("/login", spController.Login)
		userGroup.GET("/users", spController.Users)
		userGroup.POST("/users", spController.AddUser)
		userGroup.PUT("/users/:id/state/:type", spController.UpdateUserState)
		userGroup.GET("/users/:id", spController.GetUserByID)
		userGroup.PUT("/users/:id", spController.UpdateUserInfo)
		userGroup.DELETE("/users/:id", spController.DeleteUser)
		//todo
		userGroup.PUT("/users/:id/role", spController.AddUser)
	}
	return r
}
