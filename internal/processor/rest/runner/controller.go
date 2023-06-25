package runner

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest/controllers"
	controllers2 "app/internal/processor/rest/controllers"
)

type ControllersManager struct {
	storageController controllers.IStorageController
	authController    controllers.IAuthController
}

func InitControllers(manager interfaces.IManager) *ControllersManager {
	return &ControllersManager{
		storageController: controllers2.InitStorageController(manager),
		authController:    controllers2.InitAuthController(manager),
	}
}

func (cm *ControllersManager) Storage() controllers.IStorageController {
	return cm.storageController
}

func (cm *ControllersManager) Auth() controllers.IAuthController {
	return cm.authController
}
