package route

import (
	"calling-bill/internal/api"
	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine) {
	r.POST("/mobile/:username/billing", api.AddCall)
	r.GET("/mobile/:username/billing", api.GetBill)
}
