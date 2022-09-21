package router

import "github.com/go-martini/martini"

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

func (r *RouterMartini) RegisterMovieInfoHandler() {
	r.router.Group("/v1/movie_infos", func(r martini.Router) {
		r.Get("/", martini.Martini{})  // Placeholder for get all movie info handler
		r.Post("/", martini.Martini{}) // Placeholder for register movie info handler
	})
}
