package routes

import (
	v1 "routes/v1"

	"github.com/gin-gonic/gin"
)

// Init initialize routing information
func Init(r *gin.Engine) *gin.Engine {
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/rect", v1.GetRect)

	return r
}
