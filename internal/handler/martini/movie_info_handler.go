package handlermartini

import (
	"encoding/json"
	"net/http"
	"tchen/web-framework-comparison/internal/payload/request"
	"tchen/web-framework-comparison/internal/payload/response"
	"tchen/web-framework-comparison/internal/service"
)

type MovieInfoHandlerMartini struct {
	movieInfoService service.MovieInfoService
}

func NewMovieInfoMartini(movieInfoService service.MovieInfoService) MovieInfoHandlerMartini {
	return MovieInfoHandlerMartini{
		movieInfoService: movieInfoService,
	}
}

func (m MovieInfoHandlerMartini) AddMovieInfoHandler(w http.ResponseWriter, r *http.Request) {
	var movieInfoRequest request.MovieInfo

	err := json.NewDecoder(r.Body).Decode(&movieInfoRequest)
	if err != nil {
		resp := response.ReturnErrorResponse(err)
		writeRespAsJson(w, http.StatusInternalServerError, resp)
		return
	}

	movieInfoResponse, err := m.movieInfoService.AddMovieInfo(r.Context(), movieInfoRequest)
	if err != nil {
		resp := response.ReturnErrorResponse(err)
		writeRespAsJson(w, http.StatusInternalServerError, resp)
		return
	}

	resp := response.ReturnSuccessResponse(movieInfoResponse)
	writeRespAsJson(w, http.StatusOK, resp)
}

func (m MovieInfoHandlerMartini) GetAllMovieInfoHandler(w http.ResponseWriter, r *http.Request) {
	movieInfosResponse, err := m.movieInfoService.GetAllMovieInfo(r.Context())
	if err != nil {
		resp := response.ReturnErrorResponse(err)
		writeRespAsJson(w, http.StatusOK, resp)
		return
	}

	resp := response.ReturnSuccessResponse(movieInfosResponse)
	writeRespAsJson(w, http.StatusOK, resp)
}

func writeRespAsJson(w http.ResponseWriter, status int, data interface{}) {
	jsonResp, _ := json.Marshal(data)
	// w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResp)
}
