package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"myapp/dto"
	"time"
)

type PlaylistRepository interface {
	GetAllPlaylist(page, limit int) (res []dto.MsPlaylist, totalRow int, err error)
	CreatePlaylist(req *dto.MsPlaylist, createdBy string) (id string, err error)
}

type PlaylistRepositoryImpl struct {
	db *gorm.DB
}

//func (p PlaylistRepositoryImpl) CreatePlaylist(req *dto.MsPlaylist, createdBy string) (id string, err error) {
//	playlist := dto.MsPlaylist{
//		Name:      req.Name,
//		CreatedAt: time.Time{},
//		CreatedBy: "",
//		UpdatedAt: time.Time{},
//		UpdatedBy: "",
//		DeletedAt: gorm.DeletedAt{},
//		DeletedBy: "",
//	}
//	err = p.db.Create(&playlist).Error
//	log.Default().Println(err, "di service")
//	return playlist.ID, err
//}

func (p PlaylistRepositoryImpl) CreatePlaylist(req *dto.MsPlaylist, createdBy string) (id string, err error) {

	req.ID = uuid.New().String()
	req.CreatedAt = time.Now()
	req.CreatedBy = createdBy
	req.Name = req.Name

	err = p.db.Create(req).Error
	if err != nil {
		log.Println("Error saat insert:", err)
	}

	return req.ID, err
}

func (p PlaylistRepositoryImpl) GetAllPlaylist(page, limit int) (res []dto.MsPlaylist, totalRow int, err error) {
	var playlist []dto.MsPlaylist
	db := p.db.Model(&dto.MsPlaylist{})

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	log.Default().Println(db.Count(&count), "direpo")
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&playlist).Error
	return playlist, int(count), err
}

//
//func (p PlaylistRepositoryImpl) GetAllPlaylist(page, limit int) (res []dto.MsPlaylist, totalRow int, err error) {
//	var category []dto.MsPlaylist
//	db := p.db.Model(&dto.MsPlaylist{})
//
//	var count int64
//	err = db.Count(&count).Error
//	if err != nil {
//		return nil, 0, err
//	}
//
//	log.Default().Println(db.Count(&count), "direpo")
//	err = db.Offset((page - 1) * limit).Limit(limit).Find(&category).Error
//	return category, int(count), err
//}
//
//func (p PlaylistRepositoryImpl) CreatePlaylist(req *dto.MsPlaylist, createdBy string) (id string, err error) {
//	playlist := dto.MsPlaylist{
//		Name:      req.Name,
//		CreatedBy: createdBy,
//		UpdatedBy: createdBy,
//	}
//	err = p.db.Create(&playlist).Error
//	return playlist.ID, err
//}

//
//func (c CategoryRepositoryImpl) CreateRole(req *dto.MsCategoryDua, createdBy string) (id string, err error) {
//	role := dto.MsCategoryDua{
//		Name:      req.Name,
//		CreatedBy: createdBy,
//		UpdatedBy: createdBy,
//	}
//	err = c.db.Create(&role).Error
//	return role.ID, err
//}
//
//func (c CategoryRepositoryImpl) GetAllCategory(page, limit int) (res []dto.MsCategoryDua, totalRow int, err error) {
//	var category []dto.MsCategoryDua
//	db := c.db.Model(&dto.MsCategoryDua{})
//
//	var count int64
//	err = db.Count(&count).Error
//	if err != nil {
//		return nil, 0, err
//	}
//
//	log.Default().Println(db.Count(&count), "direpo")
//	err = db.Offset((page - 1) * limit).Limit(limit).Find(&category).Error
//	return category, int(count), err
//}

func NewPlaylistRepository(database *gorm.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{
		db: database,
	}
}
