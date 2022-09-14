package handlergin

import (
	"net/http"
	"tchen/web-framework-comparison/internal/payload/request"
	"tchen/web-framework-comparison/internal/payload/response"
	"tchen/web-framework-comparison/internal/service"

	"github.com/gin-gonic/gin"
)

type MovieInfoHandlerGin struct {
	movieInfoService service.MovieInfoService
}

func NewMovieInfoHandlerGin(movieInfoService service.MovieInfoService) MovieInfoHandlerGin {
	return MovieInfoHandlerGin{
		movieInfoService: movieInfoService,
	}
}

func (m MovieInfoHandlerGin) AddMovieInfoHandler(c *gin.Context) {
	var movieInfoRequest request.MovieInfo

	err := c.ShouldBindJSON(&movieInfoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ReturnErrorResponse(err))
		return
	}

	resp, err := m.movieInfoService.AddMovieInfo(c, movieInfoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ReturnErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, response.ReturnSuccessResponse(resp))

}

func (m MovieInfoHandlerGin) GetAllMovieInfoHandler(c *gin.Context) {
	resp, err := m.movieInfoService.GetAllMovieInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ReturnErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, response.ReturnSuccessResponse(resp))
}
