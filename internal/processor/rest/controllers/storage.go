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

// @Router		/v1/storage/:filename [post]
// @Tags		storage
// @Produce	json
// @Success	200	{object}	string
// @Failure	400	{object}	model.Err
// @Failure	500	{object}	model.Err
func (sc *StorageController) SaveFile(ctx *gin.Context) {
	formFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	key := ctx.Param("filename")
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

	fileName, err := sc.manager.Service().Storage().SaveFile(ctx, &model.HandlePicture{
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
