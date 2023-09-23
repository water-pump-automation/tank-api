package callbacks

import (
	"water-tank-api/controllers"

	iris "github.com/kataras/iris/v12"
)

func Get(ctx iris.Context) {
	controller := controllers.Controller{}
	tankName := ctx.Params().Get("name")

	response, err := controller.Get(tankName)

	if err != nil {
		switch response.Code {
		case controllers.NetStatNotFound:
			ctx.StatusCode(iris.StatusNotFound)
		case controllers.NetStatInternalServerError:
			ctx.StatusCode(iris.StatusInternalServerError)
		}
		ctx.JSON(response)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response)
}
