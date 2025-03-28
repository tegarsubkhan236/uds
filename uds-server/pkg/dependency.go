package pkg

import (
	"myapp/pkg/repository"
	"myapp/pkg/service"

	"gorm.io/gorm"
)

type Dependencies struct {
	PermissionService service.PermissionService
	RoleService       service.RoleService
	UserService       service.UserService
	MovieService      service.MovieService
	CategoryService   service.CategoryService
	PlaylistService   service.PlaylistService
}

func NewDependencies(db *gorm.DB) *Dependencies {
	return &Dependencies{
		PermissionService: service.NewPermissionService(repository.NewPermissionRepository(db)),
		RoleService:       service.NewRoleService(repository.NewRoleRepository(db)),
		UserService:       service.NewUserService(repository.NewUserRepository(db)),
		MovieService:      service.NewMovieService(repository.NewMovieRepository(db)),
		CategoryService:   service.NewCategoryService(repository.NewCategoryRepository(db)),
		PlaylistService:   service.NewPlaylistService(repository.NewPlaylistRepository(db)),
	}
}
