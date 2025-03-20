package dto

import (
	"gorm.io/gorm"
	"time"
)

type CrUser struct {
	ID       int    `gorm:"primary_key;auto_increment;not_null"`
	Username string `gorm:"not_null;size:255"`
	Name     string `gorm:"not_null;size:40"`
	Email    string `gorm:"not_null;size:255"`
	Password string `gorm:"not_null;size:500"`
	Status   int    `gorm:"not_null;default:2"`
	RoleID   int    `gorm:"not_null;index:idx_role_id"`
	Role     CrRole `gorm:"foreignkey:RoleID;association_foreignkey:ID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string    `gorm:"type:varchar(20);null"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string    `gorm:"type:varchar(20);null"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy string         `gorm:"type:varchar(20);null"`
}

type AuthRequest struct {
	Identity string `json:"identity" form:"identity" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserInsertRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   int    `json:"status" form:"status"`
	RoleID   int    `json:"role_id" form:"role_id" validate:"required"`
}

func (r UserInsertRequest) ToEntity() *CrUser {
	return &CrUser{
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Status:   r.Status,
		RoleID:   r.RoleID,
	}
}

type UserUpdateRequest struct {
	ID       int    `json:"id" form:"id" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   int    `json:"status" form:"status"`
	RoleID   int    `json:"role_id" form:"role_id" validate:"required"`
}

func (r UserUpdateRequest) ToEntity() *CrUser {
	return &CrUser{
		ID:       r.ID,
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Status:   r.Status,
		RoleID:   r.RoleID,
	}
}

type UserResponse struct {
	ID       int          `json:"id"`
	Username string       `json:"username"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Status   int          `json:"status"`
	Role     RoleResponse `json:"role"`
}

func (m CrUser) ToResponse() UserResponse {
	return UserResponse{
		ID:       m.ID,
		Username: m.Username,
		Name:     m.Name,
		Email:    m.Email,
		Status:   m.Status,
		Role:     m.Role.ToResponse(),
	}
}
