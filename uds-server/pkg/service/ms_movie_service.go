package service

import (
	"math"
	"mime/multipart"
	"myapp/dto"
	"myapp/pkg/repository"
	"myapp/utils"
)

type MovieService interface {
	GetMovies(page, limit int) (int, int, int, []dto.MsMovie, error)
	GetMovieById(id int) (*dto.MsMovie, error)
	CreateMovie(req *dto.MsMovie, videoFile *multipart.FileHeader, posterFile *multipart.FileHeader, createdBy string) error
	UpdateMovie(id int, req *dto.MsMovie, videoFile *multipart.FileHeader, posterFile *multipart.FileHeader, updatedBy string) error
	DeleteMovie(id int, deletedBy string) error
}

type movieServiceImpl struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieServiceImpl{
		repo: repo,
	}
}

func (m movieServiceImpl) GetMovies(page, limit int) (int, int, int, []dto.MsMovie, error) {
	result, totalRow, err := m.repo.GetMovies(page, limit)
	if err != nil {
		return 0, 0, 0, nil, err
	}

	lastPage := int(math.Ceil(float64(totalRow) / float64(limit)))

	return page, lastPage, totalRow, result, nil
}

func (m movieServiceImpl) GetMovieById(id int) (*dto.MsMovie, error) {
	return m.repo.GetMovieByID(id)
}

func (m movieServiceImpl) CreateMovie(req *dto.MsMovie, videoFile *multipart.FileHeader, posterFile *multipart.FileHeader, createdBy string) error {
	if videoFile != nil {
		videoUrl, _ := utils.Mp4ToFMp4(videoFile)
		req.VideoUrl = videoUrl
	}

	if posterFile != nil {
		posterUrl, _ := utils.SaveFile(posterFile)
		req.PosterUrl = posterUrl
	}

	if _, err := m.repo.CreateMovie(req, createdBy); err != nil {
		return err
	}

	return nil
}

func (m movieServiceImpl) UpdateMovie(id int, req *dto.MsMovie, videoFile *multipart.FileHeader, posterFile *multipart.FileHeader, updatedBy string) error {
	item, err := m.repo.GetMovieByID(id)
	if err != nil {
		return err
	}

	if videoFile != nil {
		if err := utils.RemoveFile(item.VideoUrl); err != nil {
			return err
		}
		videoUrl, _ := utils.SaveFile(videoFile)
		req.VideoUrl = videoUrl
	}

	if posterFile != nil {
		if err := utils.RemoveFile(item.PosterUrl); err != nil {
			return err
		}
		posterUrl, _ := utils.SaveFile(posterFile)
		req.PosterUrl = posterUrl
	}

	return m.repo.UpdateMovie(id, req, updatedBy)
}

func (m movieServiceImpl) DeleteMovie(id int, deletedBy string) error {
	item, err := m.repo.GetMovieByID(id)
	if err != nil {
		return err
	}

	if item.VideoUrl != "" {
		if err := utils.RemoveFile(item.VideoUrl); err != nil {
			return err
		}
	}

	if item.PosterUrl != "" {
		if err := utils.RemoveFile(item.PosterUrl); err != nil {
			return err
		}
	}

	return m.repo.DeleteMovie(id, deletedBy)
}
