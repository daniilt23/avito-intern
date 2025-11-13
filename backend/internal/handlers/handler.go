package handlers

import (
	"avito/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(Service *service.Service) *Handler {
	return &Handler{
		srv: Service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"answer": "pong"})
	})

	teamRoute := router.Group("/team")
	{
		teamRoute.POST("/add", h.CreateTeam)
		teamRoute.GET("/get", h.GetTeam)
	}

	return router
}
