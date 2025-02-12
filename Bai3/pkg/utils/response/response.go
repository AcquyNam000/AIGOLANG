package response

import (
    "github.com/kataras/iris/v12"
)

type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func Success(ctx iris.Context, data interface{}) {
    ctx.JSON(Response{
        Success: true,
        Data:    data,
    })
}

func Error(ctx iris.Context, statusCode int, message string) {
    ctx.StatusCode(statusCode)
    ctx.JSON(Response{
        Success: false,
        Error:   message,
    })
}