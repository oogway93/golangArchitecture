package handlerShopProduct

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *ProductHandler) Create(c *gin.Context) {
	var newProduct products.Product

	categoryID := c.Param("category")

	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.Create(categoryID, &newProduct)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.GetAll(categoryID)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *ProductHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	result := h.service.Get(categoryID, productID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
func (h *ProductHandler) Delete(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	result := h.service.Delete(categoryID, productID)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DELETE method doesn't work",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Category DELETED successfully",
	})
}
func (h *ProductHandler) Update(c *gin.Context) {
	var newProduct products.Product
	categoryID := c.Param("category")
	productID := c.Param("product")
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	err = h.service.Update(categoryID, productID, &newProduct)
	if err != nil {
		slog.Warn("Errors in Update handler","error", err.Error())
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category UPDATED successfully",
	})
}
