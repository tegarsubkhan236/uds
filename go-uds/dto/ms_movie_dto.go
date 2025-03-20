package dto

import (
	"gorm.io/gorm"
	"time"
)

type MsMovie struct {
	ID        int    `gorm:"primary_key"`
	Title     string `gorm:"type:varchar(40);not null"`
	VideoUrl  string `gorm:"type:varchar(255);not null"`
	PosterUrl string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string    `gorm:"type:varchar(20);null"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string    `gorm:"type:varchar(20);null"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy string         `gorm:"type:varchar(20);null"`
}

type MovieInsertRequest struct {
	Title      string `json:"title" form:"title" validate:"required"`
	VideoFile  string `json:"video_file"`
	PosterFile string `json:"poster_file"`
}

func (r *MovieInsertRequest) ToEntity() *MsMovie {
	return &MsMovie{
		Title:     r.Title,
		VideoUrl:  r.VideoFile,
		PosterUrl: r.PosterFile,
	}
}

type MovieUpdateRequest struct {
	ID        int    `json:"id" form:"id" validate:"required"`
	Title     string `json:"title" form:"title" validate:"required"`
	VideoUrl  string `json:"video_url"`
	PosterUrl string `json:"poster_url"`
}

func (r *MovieUpdateRequest) ToEntity() *MsMovie {
	return &MsMovie{
		ID:        r.ID,
		Title:     r.Title,
		VideoUrl:  r.VideoUrl,
		PosterUrl: r.PosterUrl,
	}
}

type MovieResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	VideoUrl  string `json:"video_url"`
	PosterUrl string `json:"poster_url"`
}

func (m MsMovie) ToResponse() MovieResponse {
	return MovieResponse{
		ID:        m.ID,
		Title:     m.Title,
		VideoUrl:  m.VideoUrl,
		PosterUrl: m.PosterUrl,
	}
}
