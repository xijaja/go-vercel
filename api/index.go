package handler

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Handler 是应用程序的主要入口点。可以将其视为 main() 方法
func Handler(w http.ResponseWriter, r *http.Request) {
	// 需要在 `*fiber.Ctx` 中设置正确的请求路径
	r.RequestURI = r.URL.String()
	// 通过适配器将 fiber 应用程序转换为 http.Handler
	handler().ServeHTTP(w, r)
}

// 构建 fiber 应用程序
func handler() http.HandlerFunc {
	app := fiber.New()

	app.Get("/v1", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v1",
		})
	})

	app.Get("/v2", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v2",
		})
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(), // 获取请求域名
			"path": ctx.Path(),                   // 获取请求路径
		})
	})

	return adaptor.FiberApp(app)
}
