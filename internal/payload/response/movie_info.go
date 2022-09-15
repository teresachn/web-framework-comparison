package response

import (
	"tchen/web-framework-comparison/internal/entity"
	"time"
)

type MovieInfo struct {
	ID          string    `json:"id"`
	Name        string    `json:"movie_name"`
	ReleaseDate time.Time `json:"release_date"`
}

func (m *MovieInfo) Transform(movieInfo entity.MovieInfo) {
	m.ID = movieInfo.ID.Hex()
	m.Name = movieInfo.Name
	m.ReleaseDate = movieInfo.ReleaseDate
}
