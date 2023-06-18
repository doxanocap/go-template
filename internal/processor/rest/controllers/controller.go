package controllers

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest/controllers"
)

type ControllerManager struct {
	storageController controllers.IStorageManager
}

func InitControllers(manager interfaces.IManager) *ControllerManager {
	return &ControllerManager{
		storageController: InitStorageController(manager),
	}
}

func (cm *ControllerManager) Storage() controllers.IStorageManager {
	return cm.storageController
}
