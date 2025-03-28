package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Videos struct {
	ID          string         `gorm:"type:char(36);primary_key;" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"` // Perbaiki typo dari "Tittle" ke "Title"
	Url         string         `gorm:"type:varchar(255);not null" json:"url"`
	Description string         `gorm:"type:text;null" json:"description"`
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

	// Perbaikan: One-to-Many, harus slice `[]VideoList`
	VideoList     []VideoList    `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"videoList"`
	HistoryVideo  []HistoryVideo `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"videoHistory"`
	HistoryWatch  []HistoryWatch `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"watchHistory"`
	ApprovalVideo []Approval     `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"approvalVideo"`
	Bookmark      []Bookmark     `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"bookmark"`
	Comment       []Comment      `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comment"`
}

func (m *Videos) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type VideosRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Duration    string `json:"duration" form:"duration" validate:"required"`
	Approval    string `json:"approval" form:"approval" validate:"required"`
	Privacy     string `json:"privacy" form:"privacy" validate:"required"`
	Comments    string `json:"comments" form:"comments" validate:"required"`
}

func (r VideosRequest) ToEntity() *Videos {
	playlist := Videos{
		ID:          "",
		Title:       r.Title,
		Url:         "",
		Description: r.Description,
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
		VideoList:   nil,
	}
	return &playlist
}

type VideosUpdateRequest struct {
	ID          string `json:"id" form:"id" validate:"required"`
	Tittle      string `json:"tittle" form:"tittle" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	Duration    string `json:"duration" form:"duration" validate:"required"`
	Approval    string `json:"approval" form:"approval" validate:"required"`
	Privacy     string `json:"privacy" form:"privacy" validate:"required"`
	Comments    string `json:"comments" form:"comments" validate:"required"`
}

func (r VideosUpdateRequest) ToEntity() *VideosUpdateRequest {
	playlist := VideosUpdateRequest{
		ID:          r.ID,
		Tittle:      r.Tittle,
		Description: r.Description,
		Thumbnail:   r.Thumbnail,
		Status:      r.Status,
		Duration:    r.Duration,
		Approval:    r.Approval,
		Privacy:     r.Privacy,
		Comments:    r.Comments,
	}
	return &playlist
}

type VideoResponse struct {
	Id          string `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Status      string `json:"status"`
	Duration    string `json:"duration"`
	Approval    string `json:"approval"`
	Privacy     string `json:"privacy"`
	Comments    string `json:"comments"`
}

func (m Videos) ToResponse() VideoResponse {
	videos := VideoResponse{
		Id:          m.ID,
		Tittle:      m.Title,
		Description: m.Description,
		Thumbnail:   m.Thumbnail,
		Status:      m.Status,
		Duration:    m.Duration,
		Approval:    m.Approval,
		Privacy:     m.Privacy,
		Comments:    m.Comments,
	}

	return videos
}
