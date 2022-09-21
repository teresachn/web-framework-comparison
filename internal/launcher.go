package internal

import (
	"tchen/web-framework-comparison/config"
	handlermartini "tchen/web-framework-comparison/internal/handler/martini"
	"tchen/web-framework-comparison/internal/router"
	"tchen/web-framework-comparison/internal/service"
)

func Launch() {
	var configuration config.Configuration
	configuration.StartDB()

	defer configuration.CloseDB()

	movieInfoService := service.NewMovieInfoService(configuration)
	movieInfoHandler := handlermartini.NewMovieInfoMartini(movieInfoService)

	var movieInfoRouter router.RouterMartini
	movieInfoRouter.Init()
	movieInfoRouter.RegisterMovieInfoHandler(movieInfoHandler)
	movieInfoRouter.Start()

}
