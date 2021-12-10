package hcf

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Class interface {
	Tickers(*Player) []*TickerFunc
	ArmourTiers() ArmourTiers
	Effects() []effect.Effect
	Handler(*Player) player.Handler
}

type ClassUseItem interface {
	Energy() int
	Effect() effect.Effect
	AffectMates() bool
	AffectEnemies() bool
	Item() world.Item
}

type ClassHeldItem interface {
	Effect() effect.Effect
	Item() world.Item
}

type ClassHandler struct {
	inventory.NopHandler
	P *Player
}

func (*ClassHandler) Name() string { return "ClassHandler" }

func (h *ClassHandler) HandlePlace(ctx *event.Context, slot int, i item.Stack) {
	p := h.P
	if _, ok := i.Item().(armour.Armour); ok {
		fakeContainer := *h.P.Armour().(*inventory.Armour)
		fakeContainer.Inv().AddItem(i)
		if IsClass(&fakeContainer, &Bard{}) {
			p.SetClass(NewBard(p, 120, 35))
		} else if IsClass(&fakeContainer, &Archer{}) {
			p.SetClass(NewArcher(120))
		}
	}
}
func (h *ClassHandler) HandleTake(ctx *event.Context, slot int, i item.Stack) {
	p := h.P
	if _, ok := i.Item().(armour.Armour); ok {
		p.SetClass(nil)
	}
}

func (h *ClassHandler) HandleDrop(ctx *event.Context, slot int, i item.Stack) {
	p := h.P
	if _, ok := i.Item().(armour.Armour); ok {
		p.SetClass(nil)
	}
}
