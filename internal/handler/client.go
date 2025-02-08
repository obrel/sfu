package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/obrel/sfu/internal/request"
	"github.com/obrel/sfu/internal/response"
	"github.com/obrel/sfu/internal/sfu"
)

type clientResource struct {
	sfu *sfu.SFU
}

// Routes creates a REST router for the room resource
func (rs clientResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create) // POST /clients - create a new client and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /clients/{id} - read a single client by :id
		r.Delete("/", rs.Delete) // DELETE /clients/{id}/delete - delete a single client by :id
	})

	return r
}

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client
// @Tags clients
// @Accept json
// @Produce json
// @Param client body request.CreateClientRequest true "Create client"
// @Success 200 {object} response.ClientResponse
// @Router /clients [post]
func (rs clientResource) Create(w http.ResponseWriter, r *http.Request) {
	data := &request.CreateClientRequest{}
	if err := render.Bind(r, data); err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	client := sfu.NewClient(data.Name)
	rs.sfu.Clients[client.ID()] = client
	response.CreatedWithJSON(w, &response.ClientResponse{
		ID:   client.ID(),
		Name: client.Name(),
	})
}

// GetClient godoc
// @Summary Get a client info
// @Description Create a client info
// @Tags clients
// @Accept json
// @Produce json
// @Param id path string true "Client id"
// @Success 200 {object} response.ClientResponse
// @Router /clients/{id} [get]
func (rs clientResource) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	client := rs.sfu.Clients[id]

	if client == nil {
		response.ErrNotFound(w, "Client")
		return
	}

	response.OKWithJSON(w, &response.ClientResponse{
		ID:   client.ID(),
		Name: client.Name(),
	})
}

// DeleteClient godoc
// @Summary Delete a client
// @Description Delete a client
// @Tags clients
// @Accept json
// @Produce json
// @Param id path string true "Client id"
// @Success 200
// @Router /clients/{id} [delete]
func (rs clientResource) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Clients[id] == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	delete(rs.sfu.Clients, id)
	response.OK(w)
}
