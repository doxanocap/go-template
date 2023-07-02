package runner

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest/controllers"
	controllers2 "app/internal/processor/rest/controllers"
)

type ControllersManager struct {
	authController      controllers.IAuthController
	storageController   controllers.IStorageController
	websocketController controllers.IWSController
}

func InitControllers(manager interfaces.IManager) *ControllersManager {
	return &ControllersManager{
		storageController:   controllers2.InitStorageController(manager),
		authController:      controllers2.InitAuthController(manager),
		websocketController: controllers2.InitWebsocketController(manager),
	}
}

func (cm *ControllersManager) Websocket() controllers.IWSController {
	return cm.websocketController
}

func (cm *ControllersManager) Storage() controllers.IStorageController {
	return cm.storageController
}

func (cm *ControllersManager) Auth() controllers.IAuthController {
	return cm.authController
}
