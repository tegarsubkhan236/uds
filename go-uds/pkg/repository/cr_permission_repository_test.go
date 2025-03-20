package repository

import (
	"myapp/dto"
	"myapp/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPermissionRepository(t *testing.T) {
	test.RunInTransaction(func(tx *gorm.DB) {
		repo := NewPermissionRepository(tx)

		item := &dto.CrPermission{Name: "Test Permission"}
		id, err := repo.CreatePermission(item, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, id)

		t.Run("GetPermissions", func(t *testing.T) {
			items, count, err := repo.GetPermissions(1, 10)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, count, 1)
			assert.NotEmpty(t, items)
		})

		t.Run("GetPermissionByID", func(t *testing.T) {
			result, err := repo.GetPermissionByID(id)
			assert.Nil(t, err)
			assert.Equal(t, item.Name, result.Name)
		})

		t.Run("UpdatePermission", func(t *testing.T) {
			updateReq := &dto.CrPermission{
				ID:   id,
				Name: "Updated Test Permission",
			}
			err := repo.UpdatePermission(updateReq, "admin")
			assert.Nil(t, err)

			result, _ := repo.GetPermissionByID(id)
			assert.Equal(t, updateReq.Name, result.Name)
		})

		t.Run("DeletePermission", func(t *testing.T) {
			err := repo.DeletePermission(id, "admin")
			assert.Nil(t, err)

			var deleted dto.CrPermission
			tx.Unscoped().First(&deleted, id)
			assert.NotNil(t, deleted.DeletedAt)
			assert.Equal(t, "admin", deleted.DeletedBy)
		})
	})
}
