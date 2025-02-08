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

type roomResource struct {
	sfu *sfu.SFU
}

// Routes creates a REST router for the room resource
func (rs roomResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create) // POST /rooms - create a new room and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)                 // GET /rooms/{id} - read a single room by :id
		r.Delete("/", rs.Delete)           // DELETE /rooms/{id} - delete a single room by :id
		r.Post("/join/{pid}", rs.Join)     // POST /rooms/{id}/join/{pid} - join a single participant to a room by :id
		r.Delete("/leave/{pid}", rs.Leave) // DELETE /rooms/{id}/leave/{pid} - delete a single participant from a room by :id
	})

	return r
}

// CreateRoom godoc
// @Summary Create a new room
// @Description Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body request.CreateRoomRequest true "Create room"
// @Success 200 {object} response.RoomResponse
// @Router /rooms [post]
func (rs roomResource) Create(w http.ResponseWriter, r *http.Request) {
	data := &request.CreateRoomRequest{}
	if err := render.Bind(r, data); err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	room := sfu.NewRoom(data.Name)
	rs.sfu.Rooms[room.ID()] = room
	response.CreatedWithJSON(w, &response.RoomResponse{
		ID:   room.ID(),
		Name: room.Name(),
	})
}

// GetRoom godoc
// @Summary Get a room info
// @Description Create a room info
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room id"
// @Success 200 {object} response.RoomResponse
// @Router /rooms/{id} [get]
func (rs roomResource) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	room := rs.sfu.Rooms[id]

	if room == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	var participants []*response.ParticipantPayload

	for _, p := range room.Participants() {
		participants = append(participants, &response.ParticipantPayload{
			ID:   p.ID(),
			Name: p.Name(),
		})
	}

	response.OKWithJSON(w, &response.RoomResponse{
		ID:           room.ID(),
		Name:         room.Name(),
		Participants: participants,
	})
}

// DeleteRoom godoc
// @Summary Delete a room
// @Description Delete a room
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room id"
// @Success 200
// @Router /rooms/{id} [delete]
func (rs roomResource) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Rooms[id] == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	delete(rs.sfu.Rooms, id)
	response.OK(w)
}

// JoinRoom godoc
// @Summary Join to a room
// @Description Join to a room
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room id"
// @Param pid path string true "Client id"
// @Success 200
// @Router /rooms/{id}/join/{pid} [post]
func (rs roomResource) Join(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Rooms[id] == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	room := rs.sfu.Rooms[id]

	pid, err := uuid.Parse(chi.URLParam(r, "pid"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Clients[pid] == nil {
		response.ErrNotFound(w, "Client")
		return
	}

	participant := sfu.NewParticipant(rs.sfu.Clients[pid])

	err = rs.sfu.Rooms[id].AddParticipant(participant)
	if err != nil {
		response.ErrUnprocessableEntity(w, err)
		return
	}

	var participants []*response.ParticipantPayload

	for _, p := range room.Participants() {
		participants = append(participants, &response.ParticipantPayload{
			ID:   p.ID(),
			Name: p.Name(),
		})
	}

	response.OKWithJSON(w, &response.RoomResponse{
		ID:           room.ID(),
		Name:         room.Name(),
		Participants: participants,
	})
}

// LeaveRoom godoc
// @Summary Leave to a room
// @Description Leave to a room
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room id"
// @Param pid path string true "Client id"
// @Success 200
// @Router /rooms/{id}/leave/{pid} [delete]
func (rs roomResource) Leave(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Rooms[id] == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	pid, err := uuid.Parse(chi.URLParam(r, "pid"))
	if err != nil {
		response.ErrBadRequest(w, err)
		return
	}

	if rs.sfu.Clients[pid] == nil {
		response.ErrNotFound(w, "Room")
		return
	}

	participant := sfu.NewParticipant(rs.sfu.Clients[pid])

	err = rs.sfu.Rooms[id].AddParticipant(participant)
	if err != nil {
		response.ErrUnprocessableEntity(w, err)
		return
	}

	response.OK(w)
}
