package sfu

import "github.com/google/uuid"

type SFU struct {
	Clients map[uuid.UUID]*Client
	Rooms   map[uuid.UUID]*Room
}

func NewSFU() *SFU {
	return &SFU{
		Clients: make(map[uuid.UUID]*Client),
		Rooms:   make(map[uuid.UUID]*Room),
	}
}
