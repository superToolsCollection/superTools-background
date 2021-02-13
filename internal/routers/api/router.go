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
	//r.Use(cors.Default())
	conf := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Accept", "Authorization", "X-Requested-With", "Content-Length", "Content-Type", "X-Requested-With, mytoken", "X-Requested-With, Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(conf))
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

	perManager := dao.NewSpPermissionManger("sp_permission", global.DBEngine)
	perApiManager := dao.NewSpPermissionApiManger("sp_permission_api", global.DBEngine)
	perService := service.NewSpPermissionService(perManager, perApiManager)
	perController := v1.NewSpPermissionController(perService)

	roleManager := dao.NewSpRoleManager("sp_role", global.DBEngine)
	roleService := service.NewSpRoleService(roleManager, perManager)
	roleController := v1.NewSpRoleController(roleService)

	categoryManager := dao.NewSpCategoryManager("sp_category", global.DBEngine)
	categoryService := service.NewSpCategoryService(categoryManager)
	categoryController := v1.NewSpCategoryController(categoryService)

	attributeManager := dao.NewSpAttributeManager("sp_attribute", global.DBEngine)
	attributeService := service.NewSpAttribute(attributeManager)
	attributeController := v1.NewSpAttributeController(attributeService)

	r.POST("/api/private/v1/login", spController.Login)
	userGroup := r.Group("/api/private/v1/")
	userGroup.Use(middleware.JWT())
	{
		//用户管理
		userGroup.GET("/users", spController.Users)
		userGroup.POST("/users", spController.AddUser)
		userGroup.PUT("/users/:id/state/:type", spController.UpdateUserState)
		userGroup.GET("/users/:id", spController.GetUserByID)
		userGroup.PUT("/users/:id", spController.UpdateUserInfo)
		userGroup.DELETE("/users/:id", spController.DeleteUser)
		userGroup.PUT("/users/:id/role", spController.SetRole)
		// 权限管理
		userGroup.GET("/rights/:type", perController.GetRights)
		userGroup.GET("/menus", perController.GetMenus)

		//角色相关
		userGroup.GET("/roles", roleController.GetRoleList)
		userGroup.POST("/roles", roleController.AddRole)
		userGroup.GET("/roles/:id", roleController.GetRoleById)
		userGroup.PUT("/roles/:id", roleController.UpdateRole)
		userGroup.DELETE("/roles/:id", roleController.DeleteRole)
		userGroup.POST("/roles/:id/rights", roleController.UpdateRights)
		userGroup.DELETE("/roles/:id/rights/:rightId", roleController.DeleteRight)

		//商品分类
		userGroup.GET("/categories", categoryController.GetCategoriesList)
		userGroup.POST("/categories", categoryController.AddCategory)
		userGroup.GET("/categories/:id", categoryController.GetCategory)
		userGroup.PUT("/categories/:id", categoryController.UpdateCategory)
		userGroup.DELETE("/categories/:id", categoryController.DeleteCategory)

		//分类参数管理
		userGroup.GET("categories/:id/attributes", attributeController.GetAttribute)
		userGroup.POST("categories/:id/attributes", attributeController.AddAttribute)
		userGroup.DELETE("categories/:id/attributes/:attrid", attributeController.DeleteAttribute)
		userGroup.GET("categories/:id/attributes/:attrid", attributeController.GetAttributeById)
		userGroup.PUT("categories/:id/attributes/:attrid", attributeController.UpdateAttribute)
	}
	return r
}
