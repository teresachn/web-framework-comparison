package internal

import (
	"tchen/web-framework-comparison/config"
	// 	"tchen/web-framework-comparison/internal/repository"
	handlernative "tchen/web-framework-comparison/internal/handler/native"
	"tchen/web-framework-comparison/internal/router"
	"tchen/web-framework-comparison/internal/service"
)

func Launch() {
	var configuration config.Configuration
	configuration.StartDB()

	defer configuration.CloseDB()

	// movieInfoMongoRepo := repository.NewMovieInfoRepositoryMongo(mongoConfig)
	movieInfoService := service.NewMovieInfoService(configuration)
	movieInfoHandler := handlernative.NewMovieInfoNative(movieInfoService)

	var movieInfoRouter router.RouterMux
	movieInfoRouter.Init()

	movieInfoRouter.RegisterMovieInfoHandler(movieInfoHandler)

	movieInfoRouter.Start()
}
