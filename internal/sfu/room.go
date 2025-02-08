package sfu

import (
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	mu           sync.Mutex
	id           uuid.UUID
	name         string
	participants map[uuid.UUID]*Participant
}

func NewRoom(name string) *Room {
	return &Room{
		id:           uuid.New(),
		name:         name,
		participants: make(map[uuid.UUID]*Participant),
	}
}

func (r *Room) ID() uuid.UUID {
	return r.id
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Participants() []*Participant {
	var participants []*Participant

	for _, p := range r.participants {
		participants = append(participants, p)
	}

	return participants
}

func (r *Room) GetParticipant(id uuid.UUID) *Participant {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.participants[id]
}

func (r *Room) AddParticipant(p *Participant) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.participants[p.ID()] = p
	return nil
}

func (r *Room) RemoveParticipant(p *Participant) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.participants, p.ID())
	return nil
}

func (r *Room) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return nil
}
