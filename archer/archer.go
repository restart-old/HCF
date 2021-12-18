package archer

import (
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/dragonfly-on-steroids/hcf"
)

var archerItems sync.Map

func init() {
	archerItems.Store(Sugar{}.Item(), Sugar{})
}

func (*Archer) New(p *hcf.Player) hcf.Class {
	b := &Archer{energy: 0, maxEnergy: 120}
	return b
}

type Archer struct {
	energy, maxEnergy int
	effectCoolDown    time.Time
	tickers           []*hcf.TickerFunc
}

func NewArcher(maxEnergy int) *Archer {
	return &Archer{energy: maxEnergy, maxEnergy: maxEnergy}
}

func (b *Archer) Energy() int                       { return b.energy }
func (b *Archer) MaxEnergy() int                    { return b.maxEnergy }
func (b *Archer) RemoveEnergy(energy int)           { b.energy -= energy }
func (b *Archer) HasEnoughEnergy(energy int) bool   { return b.energy >= energy }
func (b *Archer) OnEffectCoolDown() bool            { return b.effectCoolDown.After(time.Now()) }
func (b *Archer) SetEffectCoolDown(t time.Duration) { b.effectCoolDown = time.Now().Add(t) }

func (*Archer) Effects() []effect.Effect {
	return []effect.Effect{
		effect.New(effect.Speed{}, 3, 43830*time.Minute),
	}
}

func (*Archer) ArmourTiers() hcf.ArmourTiers {
	return hcf.ArmourTiers{
		Helmet:    armour.TierLeather,
		Chestlate: armour.TierLeather,
		Leggings:  armour.TierLeather,
		Boots:     armour.TierLeather,
	}
}

func (*Archer) Handler(p *hcf.Player) player.Handler {
	return &Handler{p: p}
}

type Handler struct {
	player.NopHandler
	p *hcf.Player
}

func (a *Archer) Tickers(p *hcf.Player) []*hcf.TickerFunc {
	return a.tickers
}

func (handler *Handler) HandleItemUse(ctx *event.Context) {
	player := handler.p
	if archer, ok := player.Class().(*Archer); ok {
		heldItem, _ := player.HeldItems()
		if i, ok := archerItems.Load(heldItem.Item()); ok {
			i := i.(hcf.ClassUseItem)
			if archer.OnEffectCoolDown() {
				player.Messagef("")
				return
			}
			if !archer.HasEnoughEnergy(i.Energy()) {
				player.Messagef("")
				return
			}
			player.AddEffect(i.Effect())
		}
	}
}
