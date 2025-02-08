package response

import (
	"github.com/google/uuid"
)

type ParticipantPayload struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type RoomResponse struct {
	ID           uuid.UUID             `json:"id"`
	Name         string                `json:"name"`
	Participants []*ParticipantPayload `json:"participants"`
}
