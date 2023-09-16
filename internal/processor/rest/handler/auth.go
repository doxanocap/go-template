package handler

import (
	"app/internal/cns"
	"app/internal/cns/errs"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handler) SignIn(ctx *gin.Context) {
	var body model.SignIn
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.manager.Service().User().Authenticate(ctx, body)
	if err != nil {
		if errs.IsHttpNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.SetCookie(
		"refreshToken",
		(*result).Tokens.RefreshToken,
		viper.GetInt("TOKEN_MAX_AGE"),
		viper.GetString("TOKEN_PATH"),
		viper.GetString("TOKEN_DOMAIN"),
		false,
		true,
	)

	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var body model.SignUp
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.manager.Service().User().Create(ctx, body)
	if err != nil {
		if errs.IsHttpConflictError(err) {
			ctx.Status(http.StatusConflict)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		(*result).Tokens.RefreshToken,
		viper.GetInt("TOKEN_MAX_AGE"),
		viper.GetString("TOKEN_PATH"),
		viper.GetString("TOKEN_DOMAIN"),
		false,
		true,
	)

	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) Refresh(ctx *gin.Context) {
	refreshToken := ctx.Request.Header.Get("refreshToken")

	if cns.IsNilString(refreshToken) {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	result, err := h.manager.Service().User().Refresh(ctx, refreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		(*result).RefreshToken,
		viper.GetInt("TOKEN_MAX_AGE"),
		viper.GetString("TOKEN_PATH"),
		viper.GetString("TOKEN_DOMAIN"),
		false,
		true,
	)

	ctx.JSON(http.StatusOK, *result)
}

func (h *Handler) Logout(ctx *gin.Context) {
	refreshToken := ctx.Request.Header.Get("refreshToken")

	if cns.IsNilString(refreshToken) {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if err := h.manager.Service().User().Logout(ctx, refreshToken); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		"",
		0,
		"/",
		"localhost",
		false,
		true)

	ctx.Status(http.StatusOK)
}
