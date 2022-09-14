package internal

import (
	"tchen/web-framework-comparison/config"
	handlergin "tchen/web-framework-comparison/internal/handler/handler_gin"
	"tchen/web-framework-comparison/internal/router"
	"tchen/web-framework-comparison/internal/service"
)

func Launch() {
	var configuration config.Configuration
	configuration.StartDB()

	defer configuration.CloseDB()
	movieInfoHandler := handlergin.NewMovieInfoHandlerGin(service.NewMovieInfoService(configuration))

	var router router.RouterGin
	router.Init()
	router.RegisterRouteV1(movieInfoHandler)
	router.Serve()
}
