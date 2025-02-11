package main

import (
	"golang-iris-groq/config"
	"golang-iris-groq/routes"
	"log"
	"os"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Load biến môi trường
	config.LoadEnv()

	// ✅ Sửa lỗi CORS (Dùng cách đúng của Iris v12)
	app.Use(func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		ctx.Next()
	})

	// ✅ Cấu hình API routes
	routes.SetupRoutes(app)

	// ✅ Phục vụ Frontend từ thư mục public
	app.HandleDir("/", iris.Dir("./public"))

	// ✅ Cho phép OPTIONS để tránh lỗi 405
	app.AllowMethods(iris.MethodOptions)

	// ✅ Ghi log ra console
	app.Logger().SetOutput(os.Stdout)

	// ✅ Chạy server
	port := ":8080"
	log.Printf("Server chạy tại http://localhost%s", port)
	app.Run(iris.Addr(port))
}
