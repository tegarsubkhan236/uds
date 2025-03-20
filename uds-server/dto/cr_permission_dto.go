package dto

import (
	"gorm.io/gorm"
	"time"
)

type CrPermission struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string    `gorm:"type:varchar(20);null"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string    `gorm:"type:varchar(20);null"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy string         `gorm:"type:varchar(20);null"`
}

type PermissionInsertRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (r *PermissionInsertRequest) ToEntity() *CrPermission {
	return &CrPermission{
		Name: r.Name,
	}
}

type PermissionUpdateRequest struct {
	ID   int    `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
}

func (r *PermissionUpdateRequest) ToEntity() *CrPermission {
	return &CrPermission{
		ID:   r.ID,
		Name: r.Name,
	}
}

type PermissionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m CrPermission) ToResponse() PermissionResponse {
	return PermissionResponse{
		ID:   m.ID,
		Name: m.Name,
	}
}
