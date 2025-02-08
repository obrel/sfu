package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/obrel/sfu/docs"
	"github.com/obrel/sfu/internal/sfu"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	sfu *sfu.SFU
}

func NewHandler(sfu *sfu.SFU) *Handler {
	return &Handler{
		sfu: sfu,
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (h *Handler) Service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	r.Mount("/clients", clientResource{sfu: h.sfu}.Routes())
	r.Mount("/rooms", roomResource{sfu: h.sfu}.Routes())

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4000/docs/doc.json"), //The url pointing to API definition
	))

	return r
}
