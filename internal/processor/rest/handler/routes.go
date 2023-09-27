package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
		auth.GET("/refresh", h.Refresh)
		auth.GET("/logout", h.Logout)
	}

	storage := v1.Group("/storage")
	{
		storage.POST("/:filename", h.SaveFile)
		storage.DELETE("/:id")
	}

	ws := v1.Group("/ws")
	{
		ws.GET("/pool", h.Pool)
	}
}
