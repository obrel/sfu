package response

import (
	"github.com/google/uuid"
)

type ClientResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
