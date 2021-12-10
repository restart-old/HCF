package hcf

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
)

type EffectNoLoss struct {
	new effect.Effect
	old effect.Effect
}

func (e EffectNoLoss) Add(p *player.Player) {
	p.AddEffect(e.new)
	time.AfterFunc(e.new.Duration(), func() {
		p.RemoveEffect(e.old.Type())
		p.AddEffect(e.old)
	})
}
