package handler

import (
	"github.com/andy-smoker/wh-server/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	/*pi := router.Group("/api")
	{
		wh := api.Group("/wh")
		{
			wh.POST("/")
			wh.GET("/")
			wh.GET("/:id")
			wh.PUT("/:id")
			wh.DELETE("/:id")

			item := wh.Group(":id/details")
			{
				item.GET("/")
				item.PUT("/")
			}
		}
	}*/
	return router
}