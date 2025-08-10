package services

import "github.com/kryast/crud-13.git/repositories"

type CategoryService interface{}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}
