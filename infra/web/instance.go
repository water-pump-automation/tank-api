package web

import "water-tank-api/controllers"

var instantiatedInternalController *controllers.InternalController = nil
var instantiatedExternalController *controllers.ExternalController = nil

func SetControllers(internal *controllers.InternalController, external *controllers.ExternalController) {
	instantiatedInternalController = internal
	instantiatedExternalController = external
}

func ExternalController() *controllers.ExternalController {
	return instantiatedExternalController
}

func InternalController() *controllers.InternalController {
	return instantiatedInternalController
}
