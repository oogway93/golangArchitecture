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

func (h *Handler) UserHandlerRoutes(router *gin.Engine) *gin.RouterGroup {
	user := router.Group("/user")
	{
		user.GET("/", h.GetAll)
		user.POST("/", h.Create)
		user.GET("/:login", nil)
		user.PUT("/:login", h.Update)
		user.DELETE("/:login", h.Delete)
	}

	return user
}
