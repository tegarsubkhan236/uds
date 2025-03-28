package dto

import (
	"gorm.io/gorm"
	"time"
)

type VideoList struct {
	ID         string         `gorm:"type:char(36);primary_key;" json:"id"`
	PlaylistID string         `gorm:"type:char(36);not null;index;column:playlist" json:"playlistId"`
	VideoID    string         `gorm:"type:char(36);not null;index;column:video" json:"videoId"` // Perbaikan field foreign key
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy  string         `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy  string         `gorm:"type:varchar(20);null" json:"updatedBy"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy  string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	Playlist *MsPlaylist `gorm:"foreignKey:PlaylistID;references:ID" json:"playlist"`
	Video    *Videos     `gorm:"foreignKey:VideoID;references:ID" json:"video"` // Perbaikan foreign key
}
