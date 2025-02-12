package router

import (
	"Bai3/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(dialogHandler *handlers.DialogHandler) *gin.Engine {
	r := gin.Default()

	// ✅ Chỉ tin tưởng localhost, tránh cảnh báo proxy
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.POST("/api/dialog/process", dialogHandler.ProcessDialog)

	return r
}
