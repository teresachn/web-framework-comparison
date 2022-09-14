package response

import (
	"net/http"
)

type GeneralResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func ReturnErrorResponse(err error) GeneralResponse {
	return GeneralResponse{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err.Error(),
	}
}

func ReturnBadRequestResponse(message string) GeneralResponse {
	return GeneralResponse{
		Status:  http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Error:   message,
	}
}

func ReturnSuccessResponse(data interface{}) GeneralResponse {
	return GeneralResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}
