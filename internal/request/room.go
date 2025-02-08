package request

import (
	"errors"
	"net/http"
)

type CreateRoomRequest struct {
	Name string `json:"name"`
}

func (c *CreateRoomRequest) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("Missing required fields.")
	}

	return nil
}

type JoinRoomRequest struct {
	Direction string `json:"direction"`
}

func (c *JoinRoomRequest) Bind(r *http.Request) error {
	return nil
}
