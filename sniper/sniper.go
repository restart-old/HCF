package sniper

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/dragonfly-on-steroids/hcf"
)

var DefaultSniper = &Sniper{}

type Sniper struct{}

func (*Sniper) New(p *hcf.Player) hcf.Class {
	return &Sniper{}
}

func (*Sniper) Tickers(*hcf.Player) []*hcf.TickerFunc { return nil }

func (*Sniper) ArmourTiers() hcf.ArmourTiers {
	return hcf.ArmourTiers{
		Helmet:    armour.TierChain,
		Chestlate: armour.TierLeather,
		Leggings:  armour.TierChain,
		Boots:     armour.TierLeather,
	}
}

func (*Sniper) Effects() []effect.Effect {
	return []effect.Effect{
		effect.New(effect.Speed{}, 2, 43830*time.Minute),
		effect.New(effect.Regeneration{}, 2, 43830*time.Minute),
		effect.New(effect.Resistance{}, 2, 43830*time.Minute),
	}
}

func (*Sniper) Handler(p *hcf.Player) player.Handler {
	return &Handler{p: p}
}

type Handler struct {
	player.NopHandler
	p *hcf.Player
}

func (h Handler) HandleItemUse(ctx *event.Context) {}
