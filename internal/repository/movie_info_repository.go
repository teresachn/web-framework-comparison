package repository

import (
	"context"
	entity "tchen/web-framework-comparison/internal/entity"
)

type MovieInfoRepository interface {
	AddMovieInfo(ctx context.Context, movieInfo *entity.MovieInfo) error
	GetAllMovieInfo(ctx context.Context) ([]entity.MovieInfo, error)
}
