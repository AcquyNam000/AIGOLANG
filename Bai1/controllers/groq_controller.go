package controllers

import (
	"golang-iris-groq/models"
	"golang-iris-groq/services"
	"github.com/kataras/iris/v12"
	"github.com/russross/blackfriday/v2"
)

// API gọi Groq
func HandleGroqRequest(ctx iris.Context) {
	var req models.GroqRequest

	// Kiểm tra request có dữ liệu JSON không
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu đầu vào không hợp lệ"})
		return
	}

	response, err := services.CallGroqAPI(req.Prompt)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	// Convert Markdown to HTML
	htmlResponse := string(blackfriday.Run([]byte(response)))
	ctx.JSON(iris.Map{"response": htmlResponse})
}
