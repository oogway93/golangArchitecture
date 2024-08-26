package handlerUser

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewUserHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) UserHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	user := apiRoutes.Group("/user")
	{
		user.GET("/", h.GetAll)
		user.POST("/", h.Create)
		user.GET("/:id", nil)
		user.PUT("/:login", h.Update)
		user.DELETE("/:id", nil)
	}

	return user
}
