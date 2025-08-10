package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-13.git/models"
	"github.com/kryast/crud-13.git/services"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (ch *CategoryHandler) Create(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := ch.service.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ch *CategoryHandler) GetAll(c *gin.Context) {
	category, err := ch.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ch *CategoryHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := ch.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, category)
}
