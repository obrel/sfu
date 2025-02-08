package sfu

import (
	"sync"

	"github.com/pion/webrtc/v4"
)

type PeerConnection struct {
	mu sync.Mutex
	pc *webrtc.PeerConnection
}

func NewPeerConnection(pc *webrtc.PeerConnection) *PeerConnection {
	return &PeerConnection{
		mu: sync.Mutex{},
		pc: pc,
	}
}

func (p *PeerConnection) PC() *webrtc.PeerConnection {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc
}

func (p *PeerConnection) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.Close()
}

func (p *PeerConnection) CreateOffer() (webrtc.SessionDescription, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.CreateOffer(nil)
}

func (p *PeerConnection) CreateAnswer() (webrtc.SessionDescription, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.CreateAnswer(nil)
}

func (p *PeerConnection) SetLocalDescription(sdp webrtc.SessionDescription) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.SetLocalDescription(sdp)
}

func (p *PeerConnection) SetRemoteDescription(sdp webrtc.SessionDescription) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.SetRemoteDescription(sdp)
}

func (p *PeerConnection) AddICECandidate(cand webrtc.ICECandidateInit) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.AddICECandidate(cand)
}

func (p *PeerConnection) AddTrack(track *webrtc.TrackLocalStaticRTP) (*webrtc.RTPSender, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.AddTrack(track)
}

func (p *PeerConnection) RemoveTrack(track *webrtc.RTPSender) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pc.RemoveTrack(track)
}
