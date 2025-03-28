package router

import (
	"myapp/api"
	"myapp/api/handler"
	"myapp/api/middleware"
	"myapp/pkg"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(
	r fiber.Router,
	deps *pkg.Dependencies,
) {
	r.Route("/auth", func(auth fiber.Router) {
		auth.Post("/login", handler.Login(deps.UserService))
		auth.Post("/me", middleware.Protected(), handler.Me())
		auth.Post("/logout", middleware.Protected(), handler.Logout(deps.UserService))
	})

	r.Route("/permission", func(permission fiber.Router) {
		permission.Use(middleware.Protected())
		permission.Get("/", middleware.Gateway(api.READ_PERMISSION), handler.PermissionIndex(deps.PermissionService))
		permission.Get("/show", middleware.Gateway(api.READ_PERMISSION), handler.PermissionShow(deps.PermissionService))
		permission.Post("/create", middleware.Gateway(api.CREATE_PERMISSION), handler.PermissionInsert(deps.PermissionService))
		permission.Put("/update", middleware.Gateway(api.UPDATE_PERMISSION), handler.PermissionUpdate(deps.PermissionService))
		permission.Delete("/delete", middleware.Gateway(api.DELETE_PERMISSION), handler.PermissionDelete(deps.PermissionService))
	})

	r.Route("/role", func(role fiber.Router) {
		//role.Use(middleware.Protected())
		role.Get("/", handler.RoleIndex(deps.RoleService))
		role.Get("/show", middleware.Gateway(api.READ_ROLE), handler.RoleShow(deps.RoleService))
		role.Post("/create", middleware.Gateway(api.CREATE_ROLE), handler.RoleInsert(deps.RoleService))
		role.Put("/update", middleware.Gateway(api.UPDATE_ROLE), handler.RoleUpdate(deps.RoleService))
		role.Delete("/delete", middleware.Gateway(api.DELETE_ROLE), handler.RoleDelete(deps.RoleService))
	})

	r.Route("/user", func(user fiber.Router) {
		user.Use(middleware.Protected())
		user.Get("/", handler.UserIndex(deps.UserService))
		user.Get("/show", handler.UserShow(deps.UserService))
		user.Post("/create", handler.UserInsert(deps.UserService))
		user.Put("/update", handler.UserUpdate(deps.UserService))
		user.Delete("/delete", handler.UserDelete(deps.UserService))
	})

	r.Route("/movies", func(movie fiber.Router) {
		//movie.Use(middleware.Protected())
		movie.Get("/", handler.HandleFetchAllMovie(deps.MovieService))
		movie.Get("/:id", handler.HandleFetchDetailMovie(deps.MovieService))
		movie.Post("/", handler.HandleCreateMovie(deps.MovieService))
		movie.Put("/:id", handler.HandleUpdateMovie(deps.MovieService))
		movie.Delete("/:id", handler.HandleDeleteMovie(deps.MovieService))
		movie.Get("/:id/stream", handler.HandleStreamMovie(deps.MovieService))
	})

	r.Route("/category", func(category fiber.Router) {
		//movie.Use(middleware.Protected())
		category.Get("/", handler.GetAllCategoryHandler(deps.CategoryService))
		category.Post("/create", handler.CategoryInsert(deps.CategoryService))
	})

	r.Route("/playlist", func(category fiber.Router) {
		//movie.Use(middleware.Protected())
		category.Get("/", handler.GetAllPlaylistHandler(deps.PlaylistService))
		category.Post("/create", handler.PlaylistInsert(deps.PlaylistService))
	})
}
