package routes

import (
	"golang-iris-groq/controllers"
	"github.com/kataras/iris/v12"
)

func SetupRoutes(app *iris.Application) {
	// ✅ Đăng ký API với phương thức POST
	app.Post("/api/groq", controllers.HandleGroqRequest)

	// ✅ Thêm route OPTIONS để tránh lỗi 405
	app.Options("/api/groq", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNoContent)
	})
}
