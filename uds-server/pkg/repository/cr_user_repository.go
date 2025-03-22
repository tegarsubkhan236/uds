package repository

import (
	"myapp/api"
	"myapp/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(page, limit int) (res []dto.CrUser, totalRow int, err error)
	GetUserByID(id int) (res *dto.CrUser, err error)
	GetUserByEmail(email string) (res *dto.CrUser, err error)
	CreateUser(req *dto.CrUser, createdBy string) (id int, err error)
	UpdateUser(req *dto.CrUser, updatedBy string) error
	DeleteUser(id int, deletedBy string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: database,
	}
}

func (u UserRepositoryImpl) GetAllUsers(page, limit int) (res []dto.CrUser, totalRow int, err error) {
	var users []dto.CrUser
	db := u.db.Model(&dto.CrUser{}).Where("deleted_at IS NULL").Where("status = ?", api.STATUS_ACTIVE)

	db = db.Preload("Role")
	db = db.Preload("Role.RolePermissions.Permission")

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	return users, int(count), err
}

func (u UserRepositoryImpl) GetUserByID(id int) (res *dto.CrUser, err error) {
	var data dto.CrUser
	err = u.db.Model(&dto.CrUser{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Preload("Role").
		Preload("Role.RolePermissions.Permission").
		First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (res *dto.CrUser, err error) {
	var data dto.CrUser
	err = u.db.Model(&dto.CrUser{}).
		Where("email = ? AND deleted_at IS NULL", email).
		Preload("Role").
		Preload("Role.RolePermissions.Permission").
		First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u UserRepositoryImpl) CreateUser(req *dto.CrUser, createdBy string) (id int, err error) {
	data := dto.CrUser{
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Status:    req.Status,
		RoleID:    req.RoleID,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
	err = u.db.Create(&data).Error

	return data.ID, err
}

func (u UserRepositoryImpl) UpdateUser(req *dto.CrUser, updatedBy string) error {
	updateData := map[string]any{}

	updateData["updated_by"] = updatedBy
	if req.Username != "" {
		updateData["username"] = req.Username
	}
	if req.Name != "" {
		updateData["name"] = req.Name
	}
	if req.Email != "" {
		updateData["email"] = req.Email
	}
	if req.Status != 0 && (req.Status == api.STATUS_ACTIVE || req.Status == api.STATUS_NONACTIVE) {
		updateData["status"] = req.Status
	}
	if req.RoleID != 0 {
		updateData["role_id"] = req.RoleID
	}
	if req.Password != "" {
		updateData["password"] = req.Password
	}
	if len(updateData) > 0 {
		return u.db.Model(&dto.CrUser{}).Where("id = ?", req.ID).Updates(updateData).Error
	}

	return nil
}

func (u UserRepositoryImpl) DeleteUser(id int, deletedBy string) error {
	err := u.db.Model(&dto.CrUser{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy).Error
	if err != nil {
		return err
	}

	return u.db.Delete(&dto.CrUser{}, id).Error
}
