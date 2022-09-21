package router

import (
	handlermartini "tchen/web-framework-comparison/internal/handler/martini"

	"github.com/go-martini/martini"
)

type RouterMartini struct {
	router *martini.ClassicMartini
}

func (r *RouterMartini) Init() {
	r.router = martini.Classic()
}

func (r *RouterMartini) Start() {
	// using os.Getenv("PORT") automatically
	r.router.Run()
}

func (r *RouterMartini) RegisterMovieInfoHandler(handler handlermartini.MovieInfoHandlerMartini) {
	r.router.Group("/v1/movie_infos", func(r martini.Router) {
		r.Get("/", handler.GetAllMovieInfoHandler) // Placeholder for get all movie info handler
		r.Post("/", handler.AddMovieInfoHandler)   // Placeholder for register movie info handler
	})
}
