package routers

import (
	"net/http"
	"superTools-background/internal/routers/tools"
	"time"

	_ "superTools-background/docs"
	"superTools-background/global"
	"superTools-background/internal/dao"
	"superTools-background/internal/middleware"
	"superTools-background/internal/routers/api"
	"superTools-background/internal/routers/bedtimeStory"
	"superTools-background/internal/routers/mall"
	"superTools-background/internal/routers/sd"
	"superTools-background/internal/routers/user"
	"superTools-background/internal/service"
	"superTools-background/pkg/limiter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/**
* @Author: super
* @Date: 2020-08-21 21:14
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

	//获取token
	authManager := dao.NewAuthManager("auth", global.DBEngine)
	authService := service.NewAuthService(authManager)
	authController := api.NewAuthController(authService)
	r.GET("/auth", authController.GetAuth)

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	RegisterController(r, HEALTH, global.DBEngine)
	RegisterController(r, USER, global.DBEngine)
	RegisterController(r, BEDTIME, global.DBEngine)
	RegisterController(r, PRODUCT, global.DBEngine)
	RegisterController(r, ORDER, global.DBEngine)
	RegisterController(r, TOOLS, global.DBEngine)

	return r
}

func RegisterController(r *gin.Engine, name string, db *gorm.DB) {
	switch name {
	case PRODUCT:
		registerProduct(r, db)
	case ORDER:
		registerOrder(r, db)
	case USER:
		registerUser(r, db)
	case BEDTIME:
		registerBedtime(r, db)
	case HEALTH:
		registerHealth(r, db)
	case TOOLS:
		registerTool(r, db)
	}
}

func registerHealth(r *gin.Engine, db *gorm.DB) {
	// The health check handlers
	svcd := r.Group("/sd")
	r.Use(middleware.JWT())
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
}

func registerBedtime(r *gin.Engine, db *gorm.DB) {
	storyManager := dao.NewStoryManager("stories", db)
	tagManager := dao.NewTagManager("tags", db)
	storyTagManager := dao.NewStoryTagManager("story_tag_map", db)
	storyService := service.NewStoryService(storyManager, tagManager, storyTagManager)
	storyController := bedtimeStory.NewStoryController(storyService)

	bedtime := r.Group("/api/v1/bedtime")
	bedtime.Use(middleware.JWT())
	{
		//bedtime.POST("/tags", tag.Create)
		//bedtime.DELETE("/tags/:id", tag.Delete)
		//bedtime.PUT("/tags/:id", tag.Update)
		//bedtime.PATCH("/tags/:id/state", tag.Update)
		//bedtime.GET("/tags", tag.List)
		//
		//bedtime.POST("/stories", story.Create)
		//bedtime.DELETE("/stories/:id", story.Delete)
		//bedtime.PUT("/stories/:id", story.Update)
		//bedtime.PATCH("/stories/:id/state", story.Update)
		//bedtime.GET("stories/:id", story.Get)
		//bedtime.GET("/stories", story.List)

		bedtime.GET("/story", storyController.Get)
	}
}

func registerUser(r *gin.Engine, db *gorm.DB) {
	userManager := dao.NewUserManager("users", db)
	userService := service.NewUserService(userManager)
	userController := user.NewUserController(userService)

	userGroup := r.Group("/api/v1/user/")
	{
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/register", userController.Register)
		userGroup.PUT("/update", userController.Update)
	}
}

func registerOrder(r *gin.Engine, db *gorm.DB) {
	orderManager := dao.NewOrderManager("orders", db)
	orderService := service.NewOrderService(orderManager)
	orderController := mall.NewOrderController(orderService)

	g := r.Group("/api/v1/mall")
	{
		g.GET("/orders/:id", orderController.GetOrder)
		g.GET("/all_orders", orderController.GetAllOrder)
		g.GET("/orders", orderController.GetOrderList)
		g.GET("/all_orders_user", orderController.GetOrderByUserID)
		g.GET("/orders_user", orderController.GetOrderListByUserID)
		g.POST("/orders", orderController.Insert)
		g.DELETE("/orders", orderController.Delete)
		g.PUT("/orders", orderController.Update)
	}
}

func registerProduct(r *gin.Engine, db *gorm.DB) {
	productManager := dao.NewProductManager("products", db)
	productService := service.NewProductService(productManager)
	productController := mall.NewProductController(productService)

	g := r.Group("/api/v1/mall")
	{
		g.GET("/products/:id", productController.GetProduct)
		g.GET("/all_products", productController.GetAllProduct)
		g.GET("/products", productController.GetProductList)
		g.POST("/products", productController.Insert)
		g.DELETE("/products", productController.Delete)
		g.PUT("/products", productController.Update)
	}
}

func registerTool(r *gin.Engine, db *gorm.DB) {
	toolproductManager := dao.NewToolManager("tools", db)
	toolService := service.NewToolService(toolproductManager)
	toolController := tools.NewToolController(toolService)

	g := r.Group("/api/v1/tools")
	{
		g.POST("/addTool", toolController.AddTool)
		g.PUT("/update", toolController.UpdateToolInfo)
		g.DELETE("/delete", toolController.DeleteTool)
		g.PATCH("/toolOnLine", toolController.ToolOnLine)
		g.PATCH("/toolOffLine", toolController.ToolOffLine)
		g.GET("/getTool", toolController.GetToolByKey)
		g.GET("/getToolByName", toolController.GetToolByName)
		g.GET("/toolList", toolController.GetToolList)
	}
}
