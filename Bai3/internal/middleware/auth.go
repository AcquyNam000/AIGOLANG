package middleware

import (
    "github.com/kataras/iris/v12"
)

func Auth(ctx iris.Context) {
    apiKey := ctx.GetHeader("X-API-Key")
    if apiKey == "" {
        ctx.StatusCode(iris.StatusUnauthorized)
        ctx.JSON(iris.Map{
            "error": "API key is required",
        })
        return
    }
    
    // TODO: Validate API key against configuration or database
    
    ctx.Next()
}