package routes

import (
	"github.com/MindClusters/clap-plot-api/api/user"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Clap Plot API Documents"
	docs.SwaggerInfo.Description = "API List"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = "localhost:4000"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := r.Group("/v1")
	{
		v1.POST("/login", user.Login)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
