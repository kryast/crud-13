package services

import (
	"github.com/kryast/crud-13.git/models"
	"github.com/kryast/crud-13.git/repositories"
)

type CategoryService interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
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

func (cs *categoryService) Update(category *models.Category) error {
	return cs.repo.Update(category)
}

func (cs *categoryService) Delete(id uint) error {
	return cs.repo.Delete(id)
}
