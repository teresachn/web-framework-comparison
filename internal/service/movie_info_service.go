package service

import (
	"context"
	"tchen/web-framework-comparison/internal/entity"
	request "tchen/web-framework-comparison/internal/payload/request"
	response "tchen/web-framework-comparison/internal/payload/response"
	repository "tchen/web-framework-comparison/internal/repository"
)

type MovieInfoService struct {
	repo repository.MovieInfoRepository
}

func NewMovieInfoService(repo repository.MovieInfoRepository) MovieInfoService {
	return MovieInfoService{
		repo: repo,
	}
}

func (m MovieInfoService) GetAllMovieInfo(ctx context.Context) ([]response.MovieInfo, error) {
	movieInfos, err := m.repo.GetAllMovieInfo(ctx)
	if err != nil {
		return []response.MovieInfo{}, err
	}

	var movieInfoResponses []response.MovieInfo
	for _, mI := range movieInfos {
		var movieInfoResponse response.MovieInfo
		movieInfoResponse.Transform(mI)
		movieInfoResponses = append(movieInfoResponses, movieInfoResponse)
	}

	return movieInfoResponses, nil
}

func (m MovieInfoService) AddMovieInfo(ctx context.Context, movieInfoRequest request.MovieInfo) (response.MovieInfo, error) {
	var movieInfo entity.MovieInfo
	movieInfoRequest.Transform(&movieInfo)

	err := m.repo.AddMovieInfo(ctx, &movieInfo)
	if err != nil {
		return response.MovieInfo{}, nil
	}

	var movieInfoResponse response.MovieInfo
	movieInfoResponse.Transform(movieInfo)
	return movieInfoResponse, nil
}
