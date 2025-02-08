package sfu

import "github.com/google/uuid"

type Participant struct {
	id     uuid.UUID
	client *Client
}

func NewParticipant(client *Client) *Participant {
	return &Participant{
		id:     client.id,
		client: client,
	}
}

func (p *Participant) ID() uuid.UUID {
	return p.id
}

func (p *Participant) Name() string {
	return p.client.name
}
