package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-13.git/handlers"
	"github.com/kryast/crud-13.git/repositories"
	"github.com/kryast/crud-13.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewCategoryRepository(db)
	svc := services.NewCategoryService(repo)
	h := handlers.NewCategoryHandler(svc)

	r.POST("/category", h.Create)
	r.GET("/category", h.GetAll)
	r.GET("/category/:id", h.GetByID)
	r.PUT("/category/:id", h.Update)

	return r
}
