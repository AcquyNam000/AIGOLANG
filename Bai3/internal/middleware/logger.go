package middleware

import (
    "github.com/kataras/iris/v12"
    "time"
    "log"
)

func Logger(ctx iris.Context) {
    start := time.Now()
    
    ctx.Next()
    
    // Log sau khi request được xử lý
    log.Printf(
        "%s %s %s %d %s",
        ctx.Method(),
        ctx.Path(),
        ctx.RemoteAddr(),
        ctx.GetStatusCode(),
        time.Since(start),
    )
}