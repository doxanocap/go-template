package handler

import (
	"app/internal/consts"
	"app/internal/model"
	"fmt"
	"net/http"

	"github.com/doxanocap/pkg/errs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SaveFile(ctx *gin.Context) {
	log := h.log.Named("[SaveFile]")

	formFile, err := ctx.FormFile("file")
	if err != nil {
		log.Error(fmt.Sprintf("handle file", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	key := ctx.Param("filename")
	if consts.IsNilString(key)  {
		log.Error("invalid filename")
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	file, err := formFile.Open()
	if err != nil {
		log.Error(fmt.Sprintf("open file: %s", err))
		ctx.JSON(http.StatusConflict, model.HttpConflictError)
		return
	}

	fileName, err := h.manager.Service().Storage().SaveFile(ctx, &model.HandlePicture{
		File:     file,
		FileName: formFile.Filename,
		Key:      key,
		Size:     formFile.Size,
	})

	if err != nil {
		log.Error(fmt.Sprintf("service.storage.SaveFile: %s", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusConflict {
			ctx.JSON(code, model.ErrInvalidFileFormat)
			return
		}		
		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}

	ctx.JSON(200, gin.H{"filename": fileName})
}
