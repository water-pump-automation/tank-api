package callbacks

import (
	"encoding/json"
	"water-tank-api/controllers"
	data "water-tank-api/core/entity/water_tank"
	"water-tank-api/infra/web"

	iris "github.com/kataras/iris/v12"
)

type PostBody struct {
	Name            string        `json:"name"`
	Group           string        `json:"group"`
	MaximumCapacity data.Capacity `json:"maximum_capacity"`
}

func Post(ctx iris.Context) {
	var body PostBody

	controller := web.Controller()

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
}
