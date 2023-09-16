package handler

import (
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SaveFile(ctx *gin.Context) {
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

	fileName, err := h.manager.Service().Storage().SaveFile(ctx, &model.HandlePicture{
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
