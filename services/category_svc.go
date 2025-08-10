package services

import (
	"github.com/kryast/crud-13.git/models"
	"github.com/kryast/crud-13.git/repositories"
)

type CategoryService interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (cs *categoryService) Create(category *models.Category) error {
	return cs.repo.Create(category)
}

func (cs *categoryService) GetAll() ([]models.Category, error) {
	return cs.repo.GetAll()
}

func (cs *categoryService) GetByID(id uint) (*models.Category, error) {
	return cs.repo.GetByID(id)
}
