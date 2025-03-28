package service

import (
	"log"
	"math"
	"myapp/dto"
	"myapp/pkg/repository"
)

type CategoryService interface {
	GetCategory(page, limit int) (currentPage, lastPage, totalRow int, res []dto.MsCategory, err error)
	CreateCategory(req *dto.MsCategory, createdBy string) error
	UpdateCategory(req *dto.MsCategory, updatedBy string) error
}

type CategoryServiceImpl struct {
	repo repository.CategoryRepository
}

func (r *CategoryServiceImpl) CreateCategory(req *dto.MsCategory, createdBy string) error {
	if _, err := r.repo.CreateCategory(req, createdBy); err != nil {
		return err
	}
	return nil
}

func (c CategoryServiceImpl) GetCategory(page, limit int) (currentPage, lastPage, totalRow int, res []dto.MsCategory, err error) {
	result, totalRow, err := c.repo.GetAllCategory(page, limit)
	log.Default().Println(result, "di services")
	if err != nil {
		return 0, 0, 0, nil, err
	}

	lastPage = int(math.Ceil(float64(totalRow) / float64(limit)))

	return page, lastPage, totalRow, result, nil
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		repo: repo,
	}
}
func (r *CategoryServiceImpl) UpdateCategory(req *dto.MsCategory, updatedBy string) error {
	return r.repo.UpdateCategory(req, updatedBy)
}
