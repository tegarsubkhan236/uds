package repository

import (
	"myapp/dto"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetPermissions(page, limit int) (res []*dto.CrPermission, totalRow int, err error)
	GetPermissionByID(id int) (res *dto.CrPermission, err error)
	CreatePermission(req *dto.CrPermission, createdBy string) (id int, err error)
	UpdatePermission(req *dto.CrPermission, updatedBy string) error
	DeletePermission(id int, deletedBy string) error
}

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissions(page, limit int) (res []*dto.CrPermission, totalRow int, err error) {
	var permissions []*dto.CrPermission
	db := r.db.Model(&dto.CrPermission{}).Where("deleted_at IS NULL")

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * limit).Limit(limit).Find(&permissions).Error
	return permissions, int(count), err
}

func (r *PermissionRepositoryImpl) GetPermissionByID(id int) (res *dto.CrPermission, err error) {
	permission := new(dto.CrPermission)
	err = r.db.Model(permission).Where("id = ? AND deleted_at IS NULL", id).First(permission).Error
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(req *dto.CrPermission, createdBy string) (id int, err error) {
	permission := dto.CrPermission{
		Name:      req.Name,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
	err = r.db.Create(&permission).Error
	return permission.ID, err
}

func (r *PermissionRepositoryImpl) UpdatePermission(req *dto.CrPermission, updatedBy string) error {
	data := map[string]any{
		"Name":      req.Name,
		"UpdatedBy": updatedBy,
	}
	return r.db.Model(&dto.CrPermission{}).
		Where("id = ?", req.ID).
		Updates(data).Error
}

func (r *PermissionRepositoryImpl) DeletePermission(deletedID int, deletedBy string) error {
	err := r.db.Model(&dto.CrPermission{}).
		Where("id = ?", deletedID).
		Update("deleted_by", deletedBy).Error
	if err != nil {
		return err
	}

	return r.db.Delete(&dto.CrPermission{}, deletedID).Error
}
