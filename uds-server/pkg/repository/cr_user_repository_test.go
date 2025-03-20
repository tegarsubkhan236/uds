package repository

import (
	"myapp/dto"
	"myapp/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	test.RunInTransaction(func(tx *gorm.DB) {
		roleRepo := NewRoleRepository(tx)
		role := &dto.CrRole{Name: "Test Role"}
		roleID, err := roleRepo.CreateRole(role, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, roleID)

		userRepo := NewUserRepository(tx)
		userRequest := &dto.CrUser{
			Username: "testuser",
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			Status:   1,
			RoleID:   roleID,
		}
		userID, err := userRepo.CreateUser(userRequest, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, userID)

		t.Run("GetAllUsers", func(t *testing.T) {
			users, count, err := userRepo.GetAllUsers(1, 10)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, count, 1)
			assert.NotEmpty(t, users)
		})

		t.Run("GetUserByID", func(t *testing.T) {
			user, err := userRepo.GetUserByID(userID)
			assert.Nil(t, err)
			assert.Equal(t, "testuser", user.Username)
		})

		t.Run("GetUserByEmail", func(t *testing.T) {
			user, err := userRepo.GetUserByEmail("test@example.com")
			assert.Nil(t, err)
			assert.Equal(t, "testuser", user.Username)
		})

		t.Run("UpdateUser", func(t *testing.T) {
			updateReq := &dto.CrUser{
				ID:       userID,
				Username: "updateduser",
				Name:     "Updated User",
				Email:    "updated@example.com",
			}
			err := userRepo.UpdateUser(updateReq, "admin")
			assert.Nil(t, err)

			user, _ := userRepo.GetUserByID(userID)
			assert.Equal(t, "updateduser", user.Username)
		})

		t.Run("DeleteUser", func(t *testing.T) {
			err := userRepo.DeleteUser(userID, "admin")
			assert.Nil(t, err)

			var deleted dto.CrUser
			tx.Unscoped().First(&deleted, userID)
			assert.NotNil(t, deleted.DeletedAt)
		})
	})
}
