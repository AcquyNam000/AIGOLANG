package middleware

import (
    "github.com/kataras/iris/v12"
)

func CORS(ctx iris.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    ctx.Header("Access-Control-Allow-Headers", "Content-Type, X-API-Key")
    
    if ctx.Method() == "OPTIONS" {
        ctx.StatusCode(204)
        return
    }
    
    ctx.Next()
}