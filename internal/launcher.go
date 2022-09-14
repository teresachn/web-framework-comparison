package internal

import (
	"tchen/web-framework-comparison/config"
	// 	"tchen/web-framework-comparison/internal/repository"
	// 	"tchen/web-framework-comparison/internal/service"
)

func Launch() {
	var configuration config.Configuration
	configuration.StartDB()

	defer configuration.CloseDB()

	// movieInfoMongoRepo := repository.NewMovieInfoRepositoryMongo(mongoConfig)
	// movieInfoService := service.NewMovieInfoService(movieInfoMongoRepo)

}
