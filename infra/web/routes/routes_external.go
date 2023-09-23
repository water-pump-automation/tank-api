package routes

import (
	"water-tank-api/controllers"
	"water-tank-api/infra/web/callbacks"

	iris "github.com/kataras/iris/v12"
)

type ExternalRouter struct{}

func (r *ExternalRouter) Route(i *iris.Application) {
	waterTankAPI := i.Party("/v1/water-tank")

	i.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)

		responseError := controllers.NewControllerError(controllers.WaterTankNotFound, "Route not found")
		ctx.JSON(responseError)
	})

	waterTankAPI.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{"status": "ok"})
		return
	})

	waterTankAPI.Get("/:tank", callbacks.Get)
	waterTankAPI.Get("/:group", callbacks.GetAll)

}
