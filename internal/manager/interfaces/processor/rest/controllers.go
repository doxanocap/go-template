package rest

import "app/internal/manager/interfaces/processor/rest/controllers"

type IControllersManager interface {
	Auth() controllers.IAuthController
	Storage() controllers.IStorageController
	Websocket() controllers.IWSController
}
