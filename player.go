package hcf

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/entity/physics"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Player struct {
	class Class
	*player.Player
}

func NewPlayer(p *player.Player) *Player {
	return &Player{
		class:  nil,
		Player: p,
	}
}

func HasEffectUnderLVL(p *player.Player, e2 effect.Effect, lvl int) (effect.Effect, bool) {
	for _, e := range p.Effects() {
		if e.Level() <= lvl && e.Type() == e2.Type() {
			return e, true
		}
	}
	return effect.Effect{}, false
}

func (p *Player) Class() Class { return p.class }
func (p *Player) SetClass(class Class) {
	if p.class != nil {
		for _, t := range p.class.Tickers(p) {
			go t.Stop()
		}
		for _, e := range p.class.Effects() {
			p.RemoveEffect(e.Type())
		}
	}
	var handler player.Handler
	if class != nil {
		for _, t := range class.Tickers(p) {
			go t.Start()
		}
		for _, newE := range class.Effects() {
			p.AddEffect(newE)
		}
		handler = class.Handler(p)
	}
	p.Handle(handler)
	p.class = class
}

func (p *Player) PlayersInRadius(radius float64) (players []*player.Player) {
	p1 := mgl64.Vec3{-radius, -radius, -radius}.Add(p.Position())
	p2 := mgl64.Vec3{radius, radius, radius}.Add(p.Position())

	physic := physics.NewAABB(p1, p2)
	p.World().EntitiesWithin(physic, func(e world.Entity) bool {
		p, ok := e.(*player.Player)
		if ok {
			players = append(players, p)
		}
		return false
	})
	return
}
