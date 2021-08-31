package view

import (
	"GoRestAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:list_id", h.getListById)
			lists.PATCH("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)

			items := lists.Group(":list_id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PATCH("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}

	return router
}
