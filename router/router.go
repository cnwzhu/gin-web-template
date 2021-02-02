package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"jet.com/infrared/endpoint"
)

// @title app api
// @version 1.0.o
// @description app api

// @contact.name API Support
// @contact.email me

// @BasePath /api/v1/app/
func Router(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/")
	{
		v1 := api.Group("v1/")
		{
			parking := v1.Group("app/")
			{
				parking.GET("index", endpoint.Index)
			}
		}
	}
}
