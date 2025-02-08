package sfu

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type SFU struct {
	mu      sync.Mutex
	Clients map[uuid.UUID]*Client
	Rooms   map[uuid.UUID]*Room
}

func NewSFU() *SFU {
	return &SFU{
		Clients: make(map[uuid.UUID]*Client),
		Rooms:   make(map[uuid.UUID]*Room),
	}
}

func (s *SFU) GetClient(id uuid.UUID) *Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Clients[id]
}

func (s *SFU) RegisterClient(c *Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cl := s.Clients[c.id]; cl != nil {
		return errors.New("Client already exists.")
	}

	s.Clients[c.id] = c
	return nil
}

func (s *SFU) UnregisterClient(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cl := s.Clients[id]; cl == nil {
		return errors.New("No client found.")
	}

	delete(s.Clients, id)
	return nil
}

func (s *SFU) GetRoom(id uuid.UUID) *Room {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Rooms[id]
}

func (s *SFU) RegisterRoom(r *Room) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cl := s.Rooms[r.id]; cl != nil {
		return errors.New("Room already exists.")
	}

	s.Rooms[r.id] = r
	return nil
}

func (s *SFU) UnregisterRoom(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cl := s.Rooms[id]; cl == nil {
		return errors.New("No room found.")
	}

	delete(s.Rooms, id)
	return nil
}
