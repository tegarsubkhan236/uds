package dto

import (
	"gorm.io/gorm"
	"time"
)

type CrRole struct {
	ID              int                `gorm:"primary_key"`
	Name            string             `gorm:"type:varchar(255);not null"`
	RolePermissions []CrRolePermission `gorm:"foreignkey:RoleID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string    `gorm:"type:varchar(20);null"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string    `gorm:"type:varchar(20);null"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy string         `gorm:"type:varchar(20);null"`
}

type RoleInsertRequest struct {
	Name          string `json:"name" form:"name" validate:"required"`
	PermissionIDs []int  `json:"permission_ids" form:"permissions"`
}

func (r RoleInsertRequest) ToEntity() *CrRole {
	role := CrRole{
		Name: r.Name,
	}

	var rolePermission []CrRolePermission
	for _, id := range r.PermissionIDs {
		rolePermission = append(rolePermission, CrRolePermission{PermissionID: id})
	}

	role.RolePermissions = rolePermission

	return &role
}

type RoleUpdateRequest struct {
	ID            int    `json:"id" form:"id" validate:"required"`
	Name          string `json:"name" form:"name" validate:"required"`
	PermissionIDs []int  `json:"permission_ids" form:"permissions"`
}

func (r RoleUpdateRequest) ToEntity() *CrRole {
	role := CrRole{
		ID:   r.ID,
		Name: r.Name,
	}

	var rolePermission []CrRolePermission
	for _, id := range r.PermissionIDs {
		rolePermission = append(rolePermission, CrRolePermission{PermissionID: id})
	}

	role.RolePermissions = rolePermission

	return &role
}

type RoleResponse struct {
	ID          int                      `json:"id"`
	Name        string                   `json:"name"`
	CreatedBy   string                   `json:"createdBy"`
	Permissions []RolePermissionResponse `json:"permissions"`
}

func (m CrRole) ToResponse() RoleResponse {
	role := RoleResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedBy: m.CreatedBy,
	}

	var permissions []RolePermissionResponse
	for _, p := range m.RolePermissions {
		permissions = append(permissions, p.ToResponse())
	}

	role.Permissions = permissions

	return role
}
