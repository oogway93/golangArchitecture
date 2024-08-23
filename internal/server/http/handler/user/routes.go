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

func (h *Handler) UserHandlerRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/", h.HelloWorldCon)
			user.POST("/", nil)
			user.GET("/:id", nil)
			user.PUT("/:id", nil)
			user.DELETE("/:id", nil)
		}
	}
	return r
}
