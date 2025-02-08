package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponse(w http.ResponseWriter, c int, r interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error."))
		return
	}

	w.WriteHeader(c)
	w.Write(response)
}

func OK(w http.ResponseWriter) {
	BuildResponse(w, http.StatusOK, &response{Status: "success"})
}

func OKWithJSON(w http.ResponseWriter, j interface{}) {
	BuildResponse(w, http.StatusOK, &response{Status: "success", Data: j})
}

func Created(w http.ResponseWriter) {
	BuildResponse(w, http.StatusCreated, &response{Status: "success"})
}

func CreatedWithJSON(w http.ResponseWriter, j interface{}) {
	BuildResponse(w, http.StatusCreated, &response{Status: "success", Data: j})
}

func WithMessage(w http.ResponseWriter, c int, m string) {
	BuildResponse(w, c, &response{Status: "success", Message: m})
}

func WithText(w http.ResponseWriter, c int, m string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(c)
	w.Write([]byte(m))
}
