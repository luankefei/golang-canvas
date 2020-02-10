package routes

import (
	v1 "routes/v1"

	"github.com/gin-gonic/gin"
	// _ "github.com/EDDYCJY/go-gin-example/docs"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	// "github.com/EDDYCJY/go-gin-example/middleware/jwt"
	// "github.com/EDDYCJY/go-gin-example/pkg/export"
	// "github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	// "github.com/EDDYCJY/go-gin-example/pkg/upload"
	// "github.com/EDDYCJY/go-gin-example/routers/api"
	// v1 "routes/v1"
)

// InitRouter initialize routing information
func Init(r *gin.Engine) *gin.Engine {
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/rect", v1.GetRect)

	return r
}
