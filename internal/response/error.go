package response

import (
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Status string `json:"status"`          // user-level status message
	Code   int64  `json:"code,omitempty"`  // application-specific error code
	Error  string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrNotFound(w http.ResponseWriter, r string) {
	BuildResponse(w, http.StatusNotFound, &ErrResponse{
		Status: "not found",
		Error:  fmt.Sprintf("%s not found.", r),
	})
}

func ErrBadRequest(w http.ResponseWriter, e error) {
	BuildResponse(w, http.StatusBadRequest, &ErrResponse{
		Status: "bad request",
		Error:  e.Error(),
	})
}

func ErrUnprocessableEntity(w http.ResponseWriter, e error) {
	BuildResponse(w, http.StatusUnprocessableEntity, &ErrResponse{
		Status: "unprocessable entity",
		Error:  e.Error(),
	})
}

func ErrInternalServerError(w http.ResponseWriter, e error) {
	BuildResponse(w, http.StatusInternalServerError, &ErrResponse{
		Status: "internal server error",
		Error:  e.Error(),
	})
}
