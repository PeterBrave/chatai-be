package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "chatai/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"chatai/middleware/jwt"
	"chatai/pkg/upload"
	"chatai/routers/api"
	"chatai/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())


	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv2 := r.Group("/api/v1")
	apiv2.GET("/cards", v1.GetBasicFunc)
	apiv2.GET("/cards/info", v1.GetBasicFuncInfo)
	apiv2.POST("/cards/submit", v1.PostBasicFuncInfo)
	apiv1 := r.Group("/api/v2")
	apiv1.Use(jwt.JWT())
	{
	}

	return r
}
