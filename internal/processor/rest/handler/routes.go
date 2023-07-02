package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/signIn", h.ctl.Auth().SignIn)
		auth.POST("/signUp", h.ctl.Auth().SignUp)
		auth.GET("/refresh", h.ctl.Auth().Refresh)
		auth.GET("/logout", h.ctl.Auth().Logout)
	}

	storage := v1.Group("/storage")
	{
		storage.POST("/:filename", h.ctl.Storage().SaveFile)
		storage.DELETE("/:id")
	}

	ws := v1.Group("/ws")
	{
		ws.GET("/pool", h.ctl.Websocket().Pool)
	}
}
