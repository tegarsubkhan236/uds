package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type HistorySearch struct {
	ID        string         `gorm:"type:char(36);primary_key;" json:"id"`
	Search    string         `gorm:"type:varchar(255);not null" json:"search"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"createdBy"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deletedBy"`
}

func (m *HistorySearch) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type HistorySearchRequest struct {
	Search   string `json:"search" form:"search" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
	Privacy  string `json:"privacy" form:"privacy" validate:"required"`
	Comments string `json:"comments" form:"comments" validate:"required"`
}

func (r HistorySearchRequest) ToEntity() *HistorySearch {
	search := HistorySearch{
		Status:    r.Status,
		Search:    r.Search,
		CreatedAt: time.Time{},
		CreatedBy: "",
		DeletedAt: gorm.DeletedAt{},
		DeletedBy: "",
	}
	return &search
}

type HistorySearchResponse struct {
	ID        string         `json:"id"`
	Search    string         `json:"search"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	DeletedBy string         `json:"deletedBy"`
}

func (m HistorySearch) ToResponse() HistorySearchResponse {
	search := HistorySearchResponse{
		ID:        m.ID,
		Search:    m.Search,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		DeletedAt: m.DeletedAt,
		DeletedBy: m.DeletedBy,
	}

	return search
}
