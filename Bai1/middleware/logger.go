package middleware

import (
	"log"
	"time"

	"github.com/kataras/iris/v12"
)

// Middleware ghi log request
func LoggerMiddleware(ctx iris.Context) {
	start := time.Now() // Bắt đầu đo thời gian request
	ctx.Next()          // Tiếp tục xử lý request
	duration := time.Since(start) // Tính thời gian xử lý
	log.Printf("[%s] %s %s %d %v", ctx.Method(), ctx.Path(), ctx.RemoteAddr(), ctx.GetStatusCode(), duration)
}
