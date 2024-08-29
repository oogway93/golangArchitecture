package handlerShopCategory

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewCategoryShopHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ShopCategoryHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	category := apiRoutes.Group("/categories")
	{
		category.GET("/", h.GetAll)
		category.POST("/", h.Create)
		category.GET("/:category", h.Get)
		category.PUT("/:category", h.Update)
		category.DELETE("/:category", h.Delete)
	}
	
	return category
}
