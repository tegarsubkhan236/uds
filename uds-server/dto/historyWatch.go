package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type HistoryWatch struct {
	ID      string `gorm:"type:char(36);primary_key;" json:"id"`
	VideoID string `gorm:"type:char(36);not null;index; column:video" json:"videoId"`
	//Video           string         `gorm:"type:varchar(255);not null" json:"video"`
	LastWatchedTime string         `gorm:"type:varchar(20);not null" json:"lastWatchedTime"`
	IsCompleted     string         `gorm:"type:varchar(20);not null" json:"isCompleted"`
	Status          string         `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy       string         `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy       string         `gorm:"type:varchar(20);null" json:"updatedBy"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy       string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

func (m *HistoryWatch) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type HistoryWatchRequest struct {
	Video           string `json:"video" form:"vido" validate:"required"`
	LastWatchedTime string `json:"lastWatchedTime" form:"lastWatchedTime" validate:"required"`
	IsCompleted     string `json:"isCompleted" form:"isCompleted" validate:"required"`
	Status          string `json:"status" form:"status" validate:"required"`
}

func (r HistoryWatchRequest) ToEntity() *HistoryWatch {
	playlist := HistoryWatch{
		VideoID:         r.Video,
		LastWatchedTime: r.LastWatchedTime,
		IsCompleted:     r.IsCompleted,
		Status:          r.Status,
	}
	return &playlist
}

type HistoryWatchResponse struct {
	Video           string `json:"video"`
	LastWatchedTime string `json:"lastWatchedTime"`
	IsCompleted     string `json:"isCompleted"`
	Status          string `json:"status"`
}

func (m HistoryWatch) ToResponse() HistoryWatchResponse {
	history := HistoryWatchResponse{
		Video:           m.VideoID,
		LastWatchedTime: m.LastWatchedTime,
		IsCompleted:     m.IsCompleted,
		Status:          m.Status,
	}

	return history
}
