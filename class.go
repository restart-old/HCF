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

var classes []Class

func RegisterClass(c Class) { classes = append(classes, c) }

type Class interface {
	Tickers(*Player) []*TickerFunc
	ArmourTiers() ArmourTiers
	Effects() []effect.Effect
	Handler(*Player) player.Handler
	New(*Player) Class
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
		fakeContainer := *h.P.Armour()
		fakeContainer.Inventory().AddItem(i)
		for _, class := range classes {
			if IsClass(&fakeContainer, class) {
				p.SetClass(class.New(p))
			}
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
