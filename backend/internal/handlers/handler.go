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

	teamsRoute := router.Group("/team")
	{
		teamsRoute.POST("/add", h.CreateTeam)
		teamsRoute.GET("/get", h.GetTeam)
	}
	usersRoute := router.Group("/users")
	{
		usersRoute.PATCH("/setIsActive", h.SetIsActive)
		usersRoute.GET("/getReview", h.GetReview)
	}
	prRoute := router.Group("/pullRequest")
	{
		prRoute.POST("/create", h.CreatePullRequest)
		prRoute.PATCH("/merge", h.SetMergeStatus)
		prRoute.POST("/reassign", h.ReassignReviewer)
	}

	return router
}
