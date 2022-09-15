package service

import (
	"context"
	"tchen/web-framework-comparison/config"
	"tchen/web-framework-comparison/internal/entity"
	request "tchen/web-framework-comparison/internal/payload/request"
	response "tchen/web-framework-comparison/internal/payload/response"
	repository "tchen/web-framework-comparison/internal/repository"
)

type MovieInfoService struct {
	config config.Configuration
}

func NewMovieInfoService(config config.Configuration) MovieInfoService {
	return MovieInfoService{
		config: config,
	}
}

func (m MovieInfoService) GetAllMovieInfo(ctx context.Context) ([]response.MovieInfo, error) {
	movieInfos, err := repository.GetAllMovieInfo(ctx, m.config)
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

	err := repository.AddMovieInfo(ctx, m.config, &movieInfo)
	if err != nil {
		return response.MovieInfo{}, err
	}

	var movieInfoResponse response.MovieInfo
	movieInfoResponse.Transform(movieInfo)
	return movieInfoResponse, nil
}
