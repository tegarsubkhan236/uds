package service

import (
	"math"
	"myapp/dto"
	"myapp/pkg/repository"
)

type PermissionService interface {
	GetPermissions(page, limit int) (currentPage, lastPage, totalRow int, res []*dto.CrPermission, err error)
	GetPermissionById(id int) (res *dto.CrPermission, err error)
	CreatePermission(req *dto.CrPermission, createdBy string) error
	UpdatePermission(req *dto.CrPermission, updatedBy string) error
	DeletePermission(id int, deletedBy string) error
}

type PermissionServiceImpl struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &PermissionServiceImpl{
		repo: repo,
	}
}

func (r *PermissionServiceImpl) GetPermissions(page, limit int) (currentPage, lastPage, totalRow int, res []*dto.CrPermission, err error) {
	result, totalRow, err := r.repo.GetPermissions(page, limit)
	if err != nil {
		return 0, 0, 0, nil, err
	}

	lastPage = int(math.Ceil(float64(totalRow) / float64(limit)))

	return page, lastPage, totalRow, result, nil
}

func (r *PermissionServiceImpl) GetPermissionById(id int) (*dto.CrPermission, error) {
	return r.repo.GetPermissionByID(id)
}

func (r *PermissionServiceImpl) CreatePermission(req *dto.CrPermission, createdBy string) error {
	if _, err := r.repo.CreatePermission(req, createdBy); err != nil {
		return err
	}
	return nil
}

func (r *PermissionServiceImpl) UpdatePermission(req *dto.CrPermission, updatedBy string) error {
	return r.repo.UpdatePermission(req, updatedBy)
}

func (r *PermissionServiceImpl) DeletePermission(id int, deletedBy string) error {
	return r.repo.DeletePermission(id, deletedBy)
}
