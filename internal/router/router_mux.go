package router

import (
	"log"
	"net/http"
	"os"
	handlernative "tchen/web-framework-comparison/internal/handler/native"

	"github.com/gorilla/mux"
)

var (
	PORT = os.Getenv("PORT")
)

type RouterMux struct {
	router *mux.Router
}

func (r *RouterMux) Init() {
	r.router = mux.NewRouter()
	r.router.Use(commonMiddleware)
}

func (r *RouterMux) Start() {
	log.Println("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, r.router)
}

func (r *RouterMux) RegisterMovieInfoHandler(handler handlernative.MovieInfoHandlerNative) {
	r.router.HandleFunc("/v1/movie_infos", handler.AddMovieInfoHandler).Methods("POST")
	r.router.HandleFunc("/v1/movie_infos", handler.GetAllMovieInfoHandler).Methods("GET")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
