package hcf

import (
	"sync"
)

type Session struct {
	players sync.Map
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) StorePlayer(p *Player) {
	s.players.Store(p.Player, p)
}
func (s *Session) DeletePlayer(p *Player) {
	s.players.Delete(p.Player)
}
func (s *Session) LoadPlayer(p *Player) (*Player, bool) {
	p2, _ := s.players.Load(p.Player)
	player, ok := p2.(*Player)
	return player, ok
}
