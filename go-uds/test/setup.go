package test

import (
	"log"
	"myapp/dto"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	mu   sync.Mutex
)

func init() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}

		if err := db.AutoMigrate(
			&dto.CrPermission{},
			&dto.CrRole{},
			&dto.CrRolePermission{},
			&dto.CrUser{},
			&dto.MsMovie{},
		); err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	})
}

func RunInTransaction(fn func(tx *gorm.DB)) {
	mu.Lock()
	defer mu.Unlock()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
		tx.Rollback()
	}()

	fn(tx)
}
