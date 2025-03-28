package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bookmark struct {
	ID        string         `gorm:"type:char(36);primary_key;" json:"id"`
	VideoID   string         `gorm:"type:char(36);not null;index;column:video" json:"videoId"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"createdBy"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

func (m *Bookmark) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type BookmarkRequest struct {
	Video    string `json:"video" form:"video" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
	Privacy  string `json:"privacy" form:"privacy" validate:"required"`
	Comments string `json:"comments" form:"comments" validate:"required"`
}

func (r BookmarkRequest) ToEntity() *Bookmark {
	search := Bookmark{
		VideoID:   r.Video,
		Status:    r.Status,
		CreatedAt: time.Time{},
		CreatedBy: "",
		DeletedAt: gorm.DeletedAt{},
		DeletedBy: "",
	}
	return &search
}

type BookmarkResponse struct {
	ID        string         `json:"id"`
	Video     string         `json:"video"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	DeletedBy string         `json:"deletedBy"`
}

func (m Bookmark) ToResponse() BookmarkResponse {
	search := BookmarkResponse{
		ID:        m.ID,
		Video:     m.VideoID,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		DeletedAt: m.DeletedAt,
		DeletedBy: m.DeletedBy,
	}
	return search
}
