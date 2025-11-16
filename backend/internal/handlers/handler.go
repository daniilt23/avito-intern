package handlers

import (
	"avito/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		srv: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	teamsRoute := router.Group("/team")
	{
		teamsRoute.POST("/add", h.CreateTeam)
		teamsRoute.GET("/get", h.AuthMiddleware(false), h.GetTeam)
	}
	usersRoute := router.Group("/users")
	{
		usersRoute.PATCH("/setIsActive", h.AuthMiddleware(true), h.SetIsActive)
		usersRoute.GET("/getReview", h.AuthMiddleware(false), h.GetReview)
	}
	prRoute := router.Group("/pullRequest")
	{
		prRoute.POST("/create", h.AuthMiddleware(true), h.CreatePullRequest)
		prRoute.PATCH("/merge", h.AuthMiddleware(true), h.SetMergeStatus)
		prRoute.POST("/reassign", h.AuthMiddleware(true), h.ReassignReviewer)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
