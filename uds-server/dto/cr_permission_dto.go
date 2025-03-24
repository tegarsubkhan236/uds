package dto

import (
	"gorm.io/gorm"
	"time"
)

type CrPermission struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"created_by"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy string         `gorm:"type:varchar(20);null" json:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deleted_by"`
}

type PermissionRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (r *PermissionRequest) ToEntity() *CrPermission {
	return &CrPermission{
		Name: r.Name,
	}
}
