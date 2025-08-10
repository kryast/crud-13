package repositories

import (
	"github.com/kryast/crud-13.git/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
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

func (cr *categoryRepository) GetAll() ([]models.Category, error) {
	var category []models.Category
	err := cr.db.Find(&category).Error

	return category, err
}

func (cr *categoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := cr.db.First(&category, id).Error

	return &category, err
}

func (cr *categoryRepository) Update(category *models.Category) error {
	return cr.db.Save(category).Error

}

func (cr *categoryRepository) Delete(id uint) error {
	return cr.db.Delete(&models.Category{}, id).Error
}
