package routes

import (
	"github.com/MindClusters/clap-plot-api/api/user"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", user.Login)
	}

	return r
}
