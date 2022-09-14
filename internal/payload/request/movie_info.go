package request

import (
	entity "tchen/web-framework-comparison/internal/entity"
	"time"
)

type MovieInfo struct {
	Name        string    `json:"movie_name"`
	ReleaseDate time.Time `json:"release_date"`
}

func (m MovieInfo) Transform(movieInfo *entity.MovieInfo) {
	movieInfo.Name = m.Name
	movieInfo.ReleaseDate = m.ReleaseDate
}
