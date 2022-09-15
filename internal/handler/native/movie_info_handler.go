package handlernative

import (
	"encoding/json"
	"net/http"
	"tchen/web-framework-comparison/internal/payload/request"
	"tchen/web-framework-comparison/internal/payload/response"
	"tchen/web-framework-comparison/internal/service"
)

type MovieInfoHandlerNative struct {
	movieInfoService service.MovieInfoService
}

func NewMovieInfoNative(movieInfoService service.MovieInfoService) MovieInfoHandlerNative {
	return MovieInfoHandlerNative{
		movieInfoService: movieInfoService,
	}
}

func (m MovieInfoHandlerNative) AddMovieInfoHandler(w http.ResponseWriter, r *http.Request) {
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

func (m MovieInfoHandlerNative) GetAllMovieInfoHandler(w http.ResponseWriter, r http.Request) {
	movieInfosResponse, err := m.movieInfoService.GetAllMovieInfo(r.Context())
	if err != nil {
		resp := response.ReturnErrorResponse(err)
		writeRespAsJson(w, http.StatusInternalServerError, resp)
		return
	}

	resp := response.ReturnSuccessResponse(movieInfosResponse)
	writeRespAsJson(w, http.StatusOK, resp)
}

func writeRespAsJson(w http.ResponseWriter, status int, data interface{}) {
	jsonResp, _ := json.Marshal(data)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResp)
}
