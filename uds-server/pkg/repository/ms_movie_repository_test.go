package repository

import (
	"myapp/dto"
	"myapp/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMovieRepository(t *testing.T) {
	test.RunInTransaction(func(tx *gorm.DB) {
		repo := NewMovieRepository(tx)

		movie := &dto.MsMovie{
			Title:     "Test Movie",
			VideoUrl:  "https://video.url",
			PosterUrl: "https://poster.url",
		}
		id, err := repo.CreateMovie(movie, "admin")
		assert.Nil(t, err)
		assert.NotZero(t, id)

		t.Run("GetMovies", func(t *testing.T) {
			movies, count, err := repo.GetMovies(1, 10)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, count, 1)
			assert.NotEmpty(t, movies)
		})

		t.Run("GetMovieByID", func(t *testing.T) {
			movie, err := repo.GetMovieByID(id)
			assert.Nil(t, err)
			assert.Equal(t, "Test Movie", movie.Title)
		})

		t.Run("UpdateMovie", func(t *testing.T) {
			updateReq := &dto.MsMovie{
				ID:        id,
				Title:     "Updated Movie",
				VideoUrl:  "https://updated.video.url",
				PosterUrl: "https://updated.poster.url",
			}
			err := repo.UpdateMovie(updateReq, "admin")
			assert.Nil(t, err)

			updatedMovie, _ := repo.GetMovieByID(id)
			assert.Equal(t, "Updated Movie", updatedMovie.Title)
		})

		t.Run("DeleteMovie", func(t *testing.T) {
			err := repo.DeleteMovie(id, "admin")
			assert.Nil(t, err)

			var deleted dto.MsMovie
			tx.Unscoped().First(&deleted, id)
			assert.NotNil(t, deleted)
			assert.Equal(t, "admin", deleted.DeletedBy)
		})
	})
}
