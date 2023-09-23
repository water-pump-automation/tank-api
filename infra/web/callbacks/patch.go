package callbacks

import (
	"encoding/json"
	"water-tank-api/controllers"
	"water-tank-api/core/entity/data"
	"water-tank-api/infra/web"

	iris "github.com/kataras/iris/v12"
)

type PatchBody struct {
	CurrentWaterLevel data.Capacity `json:"water_level"`
}

func Patch(ctx iris.Context) {
	var body PatchBody

	controller := web.InternalController()

	tankName := ctx.Params().Get("name")
	bodyBytes, _ := ctx.GetBody()

	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)

		responseError := controllers.NewControllerError(controllers.WaterTankBadRequest, "Bad request. Wrong type parameter")
		ctx.JSON(responseError)
		return
	}

	response, err := controller.Update(tankName, body.CurrentWaterLevel)

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
