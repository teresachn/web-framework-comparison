package router

import (
	"os"

	"github.com/gin-gonic/gin"

	handlergin "tchen/web-framework-comparison/internal/handler/handler_gin"
)

var (
	PORT = os.Getenv("PORT")
)

type RouterGin struct {
	ginEngine *gin.Engine
}

func (r *RouterGin) Init() {
	router := gin.New()

	r.ginEngine = router
}

func (r *RouterGin) Serve() {
	r.ginEngine.Run(":" + PORT)
}

func (r *RouterGin) RegisterRouteV1(handler handlergin.MovieInfoHandlerGin) {
	v1 := r.ginEngine.Group("/v1/movie_infos")
	v1.GET("/", handler.GetAllMovieInfoHandler)
	v1.POST("/", handler.AddMovieInfoHandler)
}
