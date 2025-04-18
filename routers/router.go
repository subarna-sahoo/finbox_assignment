package routers

import (
	"finbox_assignment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	featureFlagController := controllers.FeatureFlagController{}

	api := r.Group("/api")
	{
		flags := api.Group("/feature-flags")
		{
			flags.POST("", featureFlagController.Create)
		}
	}
}
