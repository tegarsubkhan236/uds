package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         string         `gorm:"type:char(36);primary_key;" json:"id"`
	Comment    string         `gorm:"type:varchar(255);not null" json:"comment"`
	RefComment string         `gorm:"type:varchar(20);not null" json:"refComment"`
	VideoID    string         `gorm:"type:char(36);not null;index;column:video" json:"videoId"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy  string         `gorm:"type:varchar(20);null" json:"createdBy"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy  string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

func (m *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

type CommentRequest struct {
	Comment    string         `json:"comment"`
	RefComment string         `json:"refComment"`
	Video      string         `json:"video"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy  string         `gorm:"type:varchar(20);null" json:"createdBy"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy  string         `gorm:"type:varchar(20);null" json:"deletedBy"`
}

func (r CommentRequest) ToEntity() *Comment {
	comment := Comment{
		Comment:    r.Comment,
		RefComment: r.RefComment,
		VideoID:    r.Video,
		CreatedAt:  time.Time{},
		CreatedBy:  "",
	}
	return &comment
}

type CommentsResponse struct {
	ID         string         `json:"id"`
	Comment    string         `json:"comment"`
	RefComment string         `json:"refComment"`
	VideoID    string         `gorm:"type:char(36);not null;index;column:video" json:"videoId"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy  string         `gorm:"type:varchar(20);null" json:"createdBy"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DeletedBy  string         `gorm:"type:varchar(20);null" json:"deletedBy"`

	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

func (m Comment) ToResponse() CommentsResponse {
	search := CommentsResponse{
		ID:        m.ID,
		VideoID:   m.VideoID,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		DeletedAt: m.DeletedAt,
		DeletedBy: m.DeletedBy,
	}
	return search
}
