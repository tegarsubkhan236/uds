package dto

type CrRolePermission struct {
	PermissionID int          `gorm:"primary_key"`
	RoleID       int          `gorm:"primary_key"`
	Role         CrRole       `gorm:"foreignkey:RoleID"`
	Permission   CrPermission `gorm:"foreignkey:PermissionID"`
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
