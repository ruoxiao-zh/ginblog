package routers

import (
	"net/http"

	_ "ginblog/docs"
	"ginblog/global"
	"ginblog/internal/middleware"
	"ginblog/internal/routers/api"
	v1 "ginblog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 引入 swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 中间件
	r.Use(middleware.Translations())

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.Use(middleware.JWT())
		{
			apiv1.POST("/tags", tag.Create)
			apiv1.DELETE("/tags/:id", tag.Delete)
			apiv1.PUT("/tags/:id", tag.Update)
			apiv1.PATCH("/tags/:id/state", tag.Update)
			apiv1.GET("/tags", tag.List)

			apiv1.POST("/articles", article.Create)
			apiv1.DELETE("/articles/:id", article.Delete)
			apiv1.PUT("/articles/:id", article.Update)
			apiv1.PATCH("/articles/:id/state", article.Update)
			apiv1.GET("/articles/:id", article.Get)
			apiv1.GET("/articles", article.List)
		}

	}

	return r
}
