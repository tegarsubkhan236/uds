package repository

import (
	"myapp/dto"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRoles(page, limit int) (res []dto.CrRole, totalRow int, err error)
	GetRoleByID(id int) (res *dto.CrRole, err error)
	CreateRole(req *dto.CrRole, createdBy string) (id int, err error)
	UpdateRole(req *dto.CrRole, updatedBy string) error
	DeleteRole(id int, deletedBy string) error
	AssignPermissionToRole(roleID, permissionID int) error
	RemovePermissionsFromRole(roleID int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(database *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: database,
	}
}

func (r *RoleRepositoryImpl) GetAllRoles(page, limit int) (res []dto.CrRole, totalRow int, err error) {
	var roles []dto.CrRole
	db := r.db.Model(&dto.CrRole{}).Where("deleted_at IS NULL")
	db = db.Preload("RolePermissions.Permission")

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * limit).Limit(limit).Find(&roles).Error
	return roles, int(count), err
}

func (r *RoleRepositoryImpl) GetRoleByID(id int) (res *dto.CrRole, err error) {
	var role dto.CrRole
	err = r.db.Model(&dto.CrRole{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Preload("RolePermissions.Permission").First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) CreateRole(req *dto.CrRole, createdBy string) (id int, err error) {
	role := dto.CrRole{
		Name:      req.Name,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
	err = r.db.Create(&role).Error
	return role.ID, err
}

func (r *RoleRepositoryImpl) UpdateRole(req *dto.CrRole, updatedBy string) error {
	data := map[string]any{
		"Name":      req.Name,
		"UpdatedBy": updatedBy,
	}
	return r.db.Model(&dto.CrRole{}).
		Where("id = ?", req.ID).
		Updates(data).Error

}

func (r *RoleRepositoryImpl) DeleteRole(id int, deletedBy string) error {
	err := r.db.Model(&dto.CrRole{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy).Error
	if err != nil {
		return err
	}

	return r.db.Delete(&dto.CrRole{}, id).Error
}

func (r *RoleRepositoryImpl) AssignPermissionToRole(roleID, permissionID int) error {
	crRolePermission := dto.CrRolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}
	return r.db.Create(&crRolePermission).Error
}

func (r *RoleRepositoryImpl) RemovePermissionsFromRole(roleID int) error {
	return r.db.Model(&dto.CrRolePermission{}).
		Where("role_id = ?", roleID).
		Delete(&dto.CrRolePermission{}).Error
}
