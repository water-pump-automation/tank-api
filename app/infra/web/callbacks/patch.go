package callbacks

import (
	"encoding/json"
	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/infra/web"

	iris "github.com/kataras/iris/v12"
)

type PatchBody struct {
	CurrentWaterLevel water_tank.Capacity `json:"water_level"`
}

func Patch(ctx iris.Context) {
	var body PatchBody

	controller := web.Controller()

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
}
