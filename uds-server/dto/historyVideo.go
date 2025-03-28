package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type HistoryVideo struct {
	ID          string         `gorm:"type:char(36);primary_key;" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"` // Perbaiki typo dari "Tittle" ke "Title"
	Description string         `gorm:"type:text;null" json:"description"`
	VideoID     string         `gorm:"type:char(36);not null;index; column:video" json:"videoId"` // Perbaikan foreign key
	Thumbnail   string         `gorm:"type:varchar(255);not null" json:"thumbnail"`
	Status      string         `gorm:"type:varchar(20);not null" json:"status"`
	Duration    string         `gorm:"type:varchar(20);not null" json:"duration"`
	Approval    string         `gorm:"type:varchar(20);not null" json:"approval"`
	Privacy     string         `gorm:"type:varchar(20);not null" json:"privacy"`
	Comments    string         `gorm:"type:varchar(20);not null" json:"comments"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy   string         `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy   string         `gorm:"type:varchar(20);null" json:"updatedBy"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy   string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	// Perbaikan: Nama relasi harus `Video`, bukan `VideoHistory`
	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

//type HistoryVideo struct {
//	ID          string         `gorm:"type:char(36);primary_key;" json:"id"`
//	Tittle      string         `gorm:"type:varchar(255);not null" json:"name"`
//	Description string         `gorm:"type:text;null" json:"description"`
//	VideoId     string         `gorm:"type:text;null;column:video" json:"videoId"`
//	Thumbnail   string         `gorm:"type:varchar(255);not null" json:"thumbnail"`
//	Status      string         `gorm:"type:varchar(20);not null" json:"status"`
//	Duration    string         `gorm:"type:varchar(20);not null" json:"duration"`
//	Approval    string         `gorm:"type:varchar(20);not null" json:"approval"`
//	Privacy     string         `gorm:"type:varchar(20);not null" json:"privacy"`
//	Comments    string         `gorm:"type:varchar(20);not null" json:"comments"`
//	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
//	CreatedBy   string         `gorm:"type:varchar(20);null" json:"createdBy"`
//	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
//	UpdatedBy   string         `gorm:"type:varchar(20);null" json:"updatedBy"`
//	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
//	DeletedBy   string         `gorm:"type:varchar(20);null" json:"deletedBy"`
//
//	VideoHistory *Videos `gorm:"foreignKey:VideoId;references:ID" json:"video"` // Perbaikan foreign key
//}

func (m *HistoryVideo) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type HistoryVideoRequest struct {
	VideoID     string `json:"video" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Duration    string `json:"duration" form:"duration" validate:"required"`
	Approval    string `json:"approval" form:"approval" validate:"required"`
	Privacy     string `json:"privacy" form:"privacy" validate:"required"`
	Comments    string `json:"comments" form:"comments" validate:"required"`
}

func (r HistoryVideoRequest) ToEntity() *HistoryVideo {
	history := HistoryVideo{
		Title:       r.Title,
		Description: r.Description,
		VideoID:     r.VideoID,
		Thumbnail:   r.Thumbnail,
		Status:      r.Status,
		Duration:    r.Duration,
		Approval:    r.Approval,
		Privacy:     r.Privacy,
		Comments:    r.Comments,
		CreatedAt:   time.Time{},
		CreatedBy:   "",
		UpdatedAt:   time.Time{},
		UpdatedBy:   "",
		DeletedAt:   gorm.DeletedAt{},
		DeletedBy:   "",
	}
	return &history
}

type HistoryVideoResponse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Video       string         `json:"video"`
	Thumbnail   string         `json:"thumbnail"`
	Status      string         `json:"status"`
	Duration    string         `json:"duration"`
	Approval    string         `json:"approval"`
	Privacy     string         `json:"privacy"`
	Comments    string         `json:"comments"`
	CreatedAt   time.Time      `json:"createdAt"`
	CreatedBy   string         `json:"createdBy"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	DeletedBy   string         `json:"deletedBy"`
}

func (m HistoryVideo) ToResponse() HistoryVideoResponse {
	search := HistoryVideoResponse{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Video:       m.VideoID,
		Thumbnail:   m.Thumbnail,
		Status:      m.Status,
		Duration:    m.Duration,
		Approval:    m.Approval,
		Privacy:     m.Privacy,
		Comments:    m.Comments,
		CreatedAt:   m.CreatedAt,
		CreatedBy:   m.CreatedBy,
		DeletedAt:   m.DeletedAt,
		DeletedBy:   m.DeletedBy,
	}

	return search
}
