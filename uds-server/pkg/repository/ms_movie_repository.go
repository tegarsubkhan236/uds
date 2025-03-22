package repository

import (
	"gorm.io/gorm"
	"myapp/dto"
)

type MovieRepository interface {
	GetMovies(page, limit int) (res []dto.MsMovie, totalRow int, err error)
	GetMovieByID(id int) (res *dto.MsMovie, err error)
	CreateMovie(req *dto.MsMovie, createdBy string) (id int, err error)
	UpdateMovie(req *dto.MsMovie, updatedBy string) error
	DeleteMovie(id int, deletedBy string) error
}

type movieRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieRepository(database *gorm.DB) MovieRepository {
	return &movieRepositoryImpl{
		db: database,
	}
}

func (m movieRepositoryImpl) GetMovies(page, limit int) (res []dto.MsMovie, totalRow int, err error) {
	var data []dto.MsMovie
	db := m.db.Model(&dto.MsMovie{})

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * limit).Limit(limit).Find(&data).Error
	return data, int(count), err
}

func (m movieRepositoryImpl) GetMovieByID(id int) (res *dto.MsMovie, err error) {
	var item dto.MsMovie
	if err = m.db.Model(&dto.MsMovie{}).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (m movieRepositoryImpl) CreateMovie(req *dto.MsMovie, createdBy string) (id int, err error) {
	item := dto.MsMovie{
		Title:     req.Title,
		VideoUrl:  req.VideoUrl,
		PosterUrl: req.PosterUrl,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
	err = m.db.Create(&item).Error
	return item.ID, err
}

func (m movieRepositoryImpl) UpdateMovie(req *dto.MsMovie, updatedBy string) error {
	data := map[string]any{
		"Title":     req.Title,
		"VideoUrl":  req.VideoUrl,
		"PosterUrl": req.PosterUrl,
		"UpdatedBy": updatedBy,
	}
	return m.db.Model(&dto.MsMovie{}).
		Where("id = ?", req.ID).
		Updates(data).Error
}

func (m movieRepositoryImpl) DeleteMovie(id int, deletedBy string) error {
	err := m.db.Model(&dto.MsMovie{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy).Error
	if err != nil {
		return err
	}

	return m.db.Delete(&dto.MsMovie{}, id).Error
}
