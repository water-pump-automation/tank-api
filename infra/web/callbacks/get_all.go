package callbacks

import (
	"water-tank-api/controllers"

	iris "github.com/kataras/iris/v12"
)

func GetAll(ctx iris.Context) {
	controller := controllers.Controller{}
	groupName := ctx.Params().Get("group")

	response, err := controller.Get(groupName)

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
