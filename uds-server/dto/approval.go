package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Approval struct {
	ID        string    `gorm:"type:char(36);primary_key;" json:"id"`
	VideoID   string    `gorm:"type:char(36);not null;index;column:video" json:"videoId"`
	Note      string    `gorm:"type:text;null" json:"note"`
	Approval  string    `gorm:"type:varchar(20);not null" json:"approval"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string    `gorm:"type:varchar(20);null" json:"createdBy"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy string    `gorm:"type:varchar(20);null" json:"updatedBy"`

	Video *Videos `gorm:"foreignKey:VideoID;references:ID" json:"video"`
}

type VideoRequest struct {
	Video    string `json:"video" form:"video" validate:"required"`
	Note     string `json:"notes" form:"note" validate:"required"`
	Approval string `json:"approval" form:"approval" validate:"required"`
}

func (m *Approval) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}

//func (r VideoRequest) ToEntity() *Videos {
//	video := Videos{}
//
//	var rolePermission []CrRolePermission
//	for _, id := range r.PermissionIDs {
//		rolePermission = append(rolePermission, CrRolePermission{PermissionID: id})
//	}
//
//	video.RolePermissions = rolePermission
//
//	return &video
//}

//type VideoUpdateRequest struct {
//	ID            int    `json:"id" form:"id" validate:"required"`
//	Name          string `json:"name" form:"name" validate:"required"`
//	PermissionIDs []int  `json:"permission_ids" form:"permissions"`
//}
//
//func (r RoleUpdateRequest) ToEntity() *CrRole {
//	role := CrRole{
//		ID:   r.ID,
//		Name: r.Name,
//	}
//
//	var rolePermission []CrRolePermission
//	for _, id := range r.PermissionIDs {
//		rolePermission = append(rolePermission, CrRolePermission{PermissionID: id})
//	}
//
//	role.RolePermissions = rolePermission
//
//	return &role
//}
//
//type RoleResponse struct {
//	ID          int                      `json:"id"`
//	Name        string                   `json:"name"`
//	CreatedBy   string                   `json:"createdBy"`
//	Permissions []RolePermissionResponse `json:"permissions"`
//}
//
//func (m CrRole) ToResponse() RoleResponse {
//	role := RoleResponse{
//		ID:        m.ID,
//		Name:      m.Name,
//		CreatedBy: m.CreatedBy,
//	}
//
//	var permissions []RolePermissionResponse
//	for _, p := range m.RolePermissions {
//		permissions = append(permissions, p.ToResponse())
//	}
//
//	role.Permissions = permissions
//
//	return role
//}
