package routes

import (
	"encoding/json"
	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/water_tank"

	iris "github.com/kataras/iris/v12"
)

type InternalRouter struct{}

type PatchBody struct {
	CurrentWaterLevel water_tank.Capacity `json:"water_level"`
}

type PostBody struct {
	Name            string              `json:"name"`
	Group           string              `json:"group"`
	MaximumCapacity water_tank.Capacity `json:"maximum_capacity"`
}

func (r *InternalRouter) Route(i *iris.Application, controller *controllers.Controller) {
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

	waterTankAPI.Post("/", func(ctx iris.Context) {
		var body PostBody

		bodyBytes, _ := ctx.GetBody()

		err := json.Unmarshal(bodyBytes, &body)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)

			responseError := controllers.NewControllerError(controllers.WaterTankBadRequest, "Bad request. Wrong type parameter")
			ctx.JSON(responseError)
			return
		}

		response, err := controller.Create(body.Name, body.Group, body.MaximumCapacity)

		if err != nil {
			switch response.Code {
			case controllers.WaterTankBadRequest:
				ctx.StatusCode(iris.StatusBadRequest)
			case controllers.WaterTankInvalidRequest:
				ctx.StatusCode(iris.StatusUnprocessableEntity)
			case controllers.WaterTankInternalServerError:
				ctx.StatusCode(iris.StatusInternalServerError)
			}
			ctx.JSON(response)
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(response)
	})
	waterTankAPI.Patch("/:tank", func(ctx iris.Context) {
		var body PatchBody

		tankName := ctx.Params().Get("tank")

		accessToken := ctx.Request().Header.Get("access_token")
		group := ctx.Request().Header.Get("group")

		bodyBytes, _ := ctx.GetBody()

		err := json.Unmarshal(bodyBytes, &body)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)

			responseError := controllers.NewControllerError(controllers.WaterTankBadRequest, "Bad request. Wrong type parameter")
			ctx.JSON(responseError)
			return
		}

		response, err := controller.Update(tankName, group, access.AccessToken(accessToken), body.CurrentWaterLevel)

		if err != nil {
			switch response.Code {
			case controllers.WaterTankBadRequest:
				ctx.StatusCode(iris.StatusBadRequest)
			case controllers.WaterTankInvalidRequest:
				ctx.StatusCode(iris.StatusUnprocessableEntity)
			case controllers.WaterTankInternalServerError:
				ctx.StatusCode(iris.StatusInternalServerError)
			}
			ctx.JSON(response)
			return
		}

		ctx.StatusCode(iris.StatusNoContent)
		ctx.JSON(response)
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
