package web

import "water-tank-api/controllers"

var instantiatedController *controllers.Controller = nil
var instantiatedExternalController *controllers.ExternalController = nil

func SetControllers(internal *controllers.Controller) {
	instantiatedController = internal
}

func Controller() *controllers.Controller {
	return instantiatedController
}
