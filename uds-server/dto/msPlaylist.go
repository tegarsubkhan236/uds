package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MsPlaylist struct {
	ID        string         `gorm:"type:char(36);primary_key;" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string         `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy string         `gorm:"type:varchar(20);null" json:"updatedBy"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	VideoList VideoList `gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"videoList"`
}

func (m *MsPlaylist) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type PlaylistInsertRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (r PlaylistInsertRequest) ToEntity() *MsPlaylist {
	playlist := MsPlaylist{
		Name: r.Name,
	}
	return &playlist
}

type PlaylistUpdateRequest struct {
	ID   string `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
}

func (r PlaylistUpdateRequest) ToEntity() *MsPlaylist {
	playlist := MsPlaylist{
		ID:   r.ID,
		Name: r.Name,
	}
	return &playlist
}

type PlaylistResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"createdBy"`
}

func (m MsPlaylist) ToResponse() PlaylistResponse {
	playlist := PlaylistResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedBy: m.CreatedBy,
	}

	return playlist
}
