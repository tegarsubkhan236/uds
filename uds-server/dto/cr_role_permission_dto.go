package dto

type CrRolePermission struct {
	PermissionID int          `gorm:"primary_key" json:"permission_id"`
	RoleID       int          `gorm:"primary_key" json:"role_id"`
	Role         CrRole       `gorm:"foreignkey:RoleID" json:"role"`
	Permission   CrPermission `gorm:"foreignkey:PermissionID" json:"permission"`
}

type RolePermissionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m CrRolePermission) ToResponse() RolePermissionResponse {
	return RolePermissionResponse{
		ID:   m.Permission.ID,
		Name: m.Permission.Name,
	}
}
