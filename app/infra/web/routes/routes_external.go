package routes

import (
	"water-tank-api/app/controllers"

	iris "github.com/kataras/iris/v12"
)

type ExternalRouter struct{}

func (r *ExternalRouter) Route(i *iris.Application, controller *controllers.Controller) {
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

	waterTankAPI.Get("/:tank", func(ctx iris.Context) {
		tankName := ctx.Params().Get("tank")

		groupName := ctx.Request().Header.Get("group")

		response, err := controller.Get(tankName, groupName)

		if err != nil {
			switch response.Code {
			case controllers.WaterTankNotFound:
				ctx.StatusCode(iris.StatusNotFound)
			case controllers.WaterTankInternalServerError:
				ctx.StatusCode(iris.StatusInternalServerError)
			}
			ctx.JSON(response)
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(response)
	})

	getAll := func(ctx iris.Context) {
		groupName := ctx.Params().Get("group")

		response, err := controller.GetGroup(groupName)

		if err != nil {
			switch response.Code {
			case controllers.WaterTankNotFound:
				ctx.StatusCode(iris.StatusNotFound)
			case controllers.WaterTankInternalServerError:
				ctx.StatusCode(iris.StatusInternalServerError)
			}
			ctx.JSON(response)
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(response)
	}
	waterTankAPI.Get("/group/:group", getAll)
	waterTankAPI.Get("/group", getAll)

}
