package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MsCategory struct {
	ID        string         `gorm:"type:char(36);primary_key;z" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy string         `gorm:"type:varchar(20);null" json:"updatedBy"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deletedBy"`
}

func (m *MsCategory) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type CategoryInsertRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

func (r CategoryInsertRequest) ToEntity() *MsCategory {
	role := MsCategory{
		Name:      r.Name,
		CreatedAt: time.Time{},
		CreatedBy: "",
		UpdatedAt: time.Time{},
		UpdatedBy: "",
		DeletedAt: gorm.DeletedAt{},
		DeletedBy: "",
	}
	return &role
}

type CategoryUpdateRequest struct {
	ID   string `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
}

func (r CategoryUpdateRequest) ToEntity() *MsCategory {
	category := MsCategory{
		ID:   r.ID,
		Name: r.Name,
	}
	return &category
}

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"createdBy"`
	CreatedAt string `json:"createdAt"`
}

func (m MsCategory) ToResponse() CategoryResponse {
	category := CategoryResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedBy: m.CreatedBy,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
	}
	return category
}
