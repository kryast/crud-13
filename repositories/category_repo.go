package repositories

import (
	"github.com/kryast/crud-13.git/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) Create(category *models.Category) error {
	return cr.db.Create(category).Error
}
