package controllers

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StorageController struct {
	manager interfaces.IManager
}

func InitStorageController(manager interfaces.IManager) *StorageController {
	return &StorageController{
		manager: manager,
	}
}

func (sc *StorageController) HandlePicture(ctx *gin.Context) {
	formFile, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileName, err := sc.manager.Service().Storage().HandlePicture(ctx, &model.HandlePictureReq{
		File:     file,
		FileName: formFile.Filename,
		Key:      key,
		Size:     formFile.Size,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"filename": fileName})
}
