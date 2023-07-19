package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

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
