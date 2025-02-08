package request

import (
	"errors"
	"net/http"
)

type CreateClientRequest struct {
	Name string `json:"name"`
}

func (c *CreateClientRequest) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("Missing required fields.")
	}

	return nil
}
