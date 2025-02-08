package sfu

import (
	"github.com/google/uuid"
	"github.com/pion/webrtc/v4"
)

type Client struct {
	id             uuid.UUID
	name           string
	peerConnection *PeerConnection
}

func NewClient(name string) *Client {
	return &Client{
		id:   uuid.New(),
		name: name,
	}
}

func (c *Client) ID() uuid.UUID {
	return c.id
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) PeerConnection() *PeerConnection {
	return c.peerConnection
}

func (c *Client) CreatePeerConnection(config webrtc.Configuration) error {
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return err
	}

	c.peerConnection = NewPeerConnection(pc)
	return nil
}
