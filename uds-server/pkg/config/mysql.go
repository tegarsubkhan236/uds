package config

import (
	"fmt"
	"myapp/api"
	"myapp/dto"
	"myapp/pkg/repository"
	"myapp/pkg/service"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() (*gorm.DB, error) {
	username := viper.GetString(DbUser)
	password := viper.GetString(DbPassword)
	host := viper.GetString(DbHost)
	port := viper.GetInt(DbPort)
	dbname := viper.GetString(DbName)
	parsetime := viper.GetBool(DbParsetime)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", username, password, host, port, dbname, parsetime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&dto.CrPermission{},
		&dto.CrRole{},
		&dto.CrRolePermission{},
		&dto.CrUser{},
		&dto.MsMovie{},
		&dto.MsCategory{},
		&dto.MsPlaylist{},
		&dto.Videos{},
		&dto.HistorySearch{},
		&dto.Approval{},
		&dto.HistoryWatch{},
		&dto.Comment{},
		&dto.Bookmark{},
		&dto.VideoList{},
		&dto.HistoryVideo{},
	); err != nil {
		return nil, err
	}

	if err := seedDB(db); err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the db.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func seedDB(db *gorm.DB) error {
	permissions := []dto.CrPermission{
		{Name: api.CREATE_PERMISSION, CreatedBy: "SYSTEM"},
		{Name: api.READ_PERMISSION, CreatedBy: "SYSTEM"},
		{Name: api.UPDATE_PERMISSION, CreatedBy: "SYSTEM"},
		{Name: api.DELETE_PERMISSION, CreatedBy: "SYSTEM"},
		{Name: api.CREATE_ROLE, CreatedBy: "SYSTEM"},
		{Name: api.READ_ROLE, CreatedBy: "SYSTEM"},
		{Name: api.UPDATE_ROLE, CreatedBy: "SYSTEM"},
		{Name: api.DELETE_ROLE, CreatedBy: "SYSTEM"},
		{Name: api.CREATE_USER, CreatedBy: "SYSTEM"},
		{Name: api.READ_USER, CreatedBy: "SYSTEM"},
		{Name: api.UPDATE_USER, CreatedBy: "SYSTEM"},
		{Name: api.DELETE_USER, CreatedBy: "SYSTEM"},
	}

	for _, permission := range permissions {
		if err := db.FirstOrCreate(&permission, dto.CrPermission{Name: permission.Name}).Error; err != nil {
			return err
		}
	}

	role := dto.CrRole{Name: "SUPER_ADMIN", CreatedBy: "SYSTEM"}
	if err := db.FirstOrCreate(&role, dto.CrRole{Name: role.Name}).Error; err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	superAdmin := dto.CrUser{
		Username: "admin",
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: "admin",
		RoleID:   role.ID,
		Status:   api.STATUS_ACTIVE,
	}

	if err := userService.CreateUser(&superAdmin, "SYSTEM"); err != nil {
		return err
	}

	return nil
}
