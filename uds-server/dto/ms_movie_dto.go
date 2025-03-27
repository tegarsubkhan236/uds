package dto

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MsMovie struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Title     string         `gorm:"type:varchar(40);not null" json:"title"`
	VideoUrl  string         `gorm:"type:varchar(255);not null" json:"video_url"`
	PosterUrl string         `gorm:"type:varchar(255);not null" json:"poster_url"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"created_by"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy string         `gorm:"type:varchar(20);null" json:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deleted_by"`
}

type MovieRequest struct {
	Title      string `json:"title" form:"title" validate:"required"`
	VideoFile  string `json:"video_file"`
	PosterFile string `json:"poster_file"`
}

func (r *MovieRequest) ToEntity() *MsMovie {
	return &MsMovie{
		Title:     r.Title,
		VideoUrl:  r.VideoFile,
		PosterUrl: r.PosterFile,
	}
}

type MovieResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	VideoUrl  string `json:"video_url"`
	PosterUrl string `json:"poster_url"`
}

func (m MsMovie) ToResponse() MovieResponse {
	address := fmt.Sprintf("%s:%d", viper.GetString("apps.host"), viper.GetInt("apps.port"))

	videoUrl := m.VideoUrl
	if !strings.HasPrefix(videoUrl, "http") {
		videoUrl = fmt.Sprintf("http://%s%s", address, videoUrl)
	}

	posterUrl := m.PosterUrl
	if !strings.HasPrefix(posterUrl, "http") {
		posterUrl = fmt.Sprintf("http://%s%s", address, posterUrl)
	}

	return MovieResponse{
		ID:        m.ID,
		Title:     m.Title,
		VideoUrl:  videoUrl,
		PosterUrl: posterUrl,
	}
}
