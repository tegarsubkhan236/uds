package repository

import (
	"myapp/dto"
	"myapp/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRoleRepository(t *testing.T) {
	test.RunInTransaction(func(tx *gorm.DB) {
		repo := NewRoleRepository(tx)
		permissionRepo := NewPermissionRepository(tx)

		role := &dto.CrRole{Name: "Test Role"}
		roleID, err := repo.CreateRole(role, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, roleID)

		permission := &dto.CrPermission{Name: "Test Permission"}
		permissionID, err := permissionRepo.CreatePermission(permission, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, permissionID)

		t.Run("AssignPermissionToRole", func(t *testing.T) {
			err := repo.AssignPermissionToRole(roleID, permissionID)
			assert.Nil(t, err)
		})

		t.Run("RemovePermissionsFromRole", func(t *testing.T) {
			err := repo.RemovePermissionsFromRole(roleID)
			assert.Nil(t, err)
		})

		t.Run("GetRoles", func(t *testing.T) {
			roles, count, err := repo.GetAllRoles(1, 10)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, count, 1)
			assert.NotEmpty(t, roles)
		})

		t.Run("GetRoleByID", func(t *testing.T) {
			result, err := repo.GetRoleByID(roleID)
			assert.Nil(t, err)
			assert.Equal(t, "Test Role", result.Name)
		})

		t.Run("UpdateRole", func(t *testing.T) {
			updateReq := &dto.CrRole{
				ID:   roleID,
				Name: "Updated Test Role",
			}
			err := repo.UpdateRole(updateReq, "admin")
			assert.Nil(t, err)

			result, _ := repo.GetRoleByID(roleID)
			assert.Equal(t, "Updated Test Role", result.Name)
		})

		t.Run("DeleteRole", func(t *testing.T) {
			err := repo.DeleteRole(roleID, "admin")
			assert.Nil(t, err)

			var deleted dto.CrRole
			tx.Unscoped().First(&deleted, roleID)
			assert.NotNil(t, deleted.DeletedAt)
			assert.Equal(t, "admin", deleted.DeletedBy)
		})
	})
}
