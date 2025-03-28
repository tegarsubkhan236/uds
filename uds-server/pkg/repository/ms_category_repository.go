package repository

import (
	"github.com/google/uuid"
	"log"
	"myapp/dto"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(page, limit int) (res []dto.MsCategory, totalRow int, err error)
	CreateCategory(req *dto.MsCategory, createdBy string) (id string, err error)
	UpdateCategory(req *dto.MsCategory, updatedBy string) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (c CategoryRepositoryImpl) UpdateCategory(req *dto.MsCategory, updatedBy string) error {
	data := map[string]any{
		"Name":      req.Name,
		"UpdatedBy": updatedBy,
	}
	return c.db.Model(&dto.CrPermission{}).
		Where("id = ?", req.ID).
		Updates(data).Error
}

func (c CategoryRepositoryImpl) CreateCategory(req *dto.MsCategory, createdBy string) (id string, err error) {
	req.ID = uuid.New().String()
	req.CreatedAt = time.Now()
	req.CreatedBy = createdBy
	req.Name = req.Name

	err = c.db.Create(req).Error
	if err != nil {
		log.Println("Error saat insert:", err)
	}

	return req.ID, err
}

func (c CategoryRepositoryImpl) CreateRole(req *dto.MsCategory, createdBy string) (id string, err error) {

	category := dto.MsCategory{
		Name:      req.Name,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}

	err = c.db.Create(&category).Error

	return category.ID, err
}

func (c CategoryRepositoryImpl) GetAllCategory(page, limit int) (res []dto.MsCategory, totalRow int, err error) {
	var category []dto.MsCategory
	db := c.db.Model(&dto.MsCategory{})

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	log.Default().Println(db.Count(&count), "direpo")
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&category).Error
	return category, int(count), err
}

func NewCategoryRepository(database *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: database,
	}
}
