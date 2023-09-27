package handler

import (
	"app/internal/consts"
	"app/internal/model"
	"fmt"
	"net/http"

	"github.com/doxanocap/pkg/errs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(ctx *gin.Context) {
	log := h.log.Named("[SignIn]")

	var body model.SignIn
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.manager.Service().User().Authenticate(ctx, body)
	if err != nil {
		log.Error(fmt.Sprintf("service.user.Authenticate: %s", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrUserNotFound)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
	}

	ctx.SetCookie("refreshToken", result.Tokens.RefreshToken, h.cfg.Token.MaxAge, h.cfg.Token.Path, h.cfg.Token.Domain, false, true)
	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) SignUp(ctx *gin.Context) {
	log := h.log.Named("[SignUp]")

	var body model.SignUp
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Error(fmt.Sprintf("bindJSON", err))

		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	result, err := h.manager.Service().User().Create(ctx, body)
	if err != nil {
		log.Error(fmt.Sprintf("service.user.Create: %s", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusConflict {
			ctx.JSON(code, model.ErrUserAlreadyExist)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}

	ctx.SetCookie("refreshToken", result.Tokens.RefreshToken, h.cfg.Token.MaxAge, h.cfg.Token.Path, h.cfg.Token.Domain, false, true)
	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) Refresh(ctx *gin.Context) {
	log := h.log.Named("[Refresh]")

	refreshToken := ctx.Request.Header.Get("refreshToken")

	if consts.IsNilString(refreshToken) {
		log.Error("unauthorized")
		ctx.JSON(http.StatusUnauthorized, model.HttpUnauthorized)
		return
	}

	result, err := h.manager.Service().User().Refresh(ctx, refreshToken)
	if err != nil {
		log.Error(fmt.Sprintf("service.user.Refresh: %s", err))
		code := errs.UnmarshalCode(err)
		if code == http.StatusConflict {
			ctx.JSON(code, model.ErrInvalidToken)
			return
		}
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrTokenNotFound)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}

	ctx.SetCookie("refreshToken", result.RefreshToken, h.cfg.Token.MaxAge, h.cfg.Token.Path, h.cfg.Token.Domain, false, true)
	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) Logout(ctx *gin.Context) {
	log := h.log.Named("[Logout]")

	refreshToken := ctx.Request.Header.Get("refreshToken")
	if consts.IsNilString(refreshToken) {
		log.Error("unauthorized")
		ctx.JSON(http.StatusUnauthorized, model.HttpUnauthorized)
		return
	}

	if err := h.manager.Service().User().Logout(ctx, refreshToken); err != nil {
		log.Error(fmt.Sprintf("service.user.Logout: %s", err))
		code := errs.UnmarshalCode(err)
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrTokenNotFound)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}

	ctx.SetCookie("refreshToken", "", 0, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
