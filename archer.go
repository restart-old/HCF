package hcf

import (
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
)

var archerItems sync.Map

func init() {

}

type Archer struct {
	energy, maxEnergy int
	effectCoolDown    time.Time
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

func (*Archer) ArmourTiers() ArmourTiers {
	return ArmourTiers{
		Helmet:    armour.TierLeather,
		Chestlate: armour.TierLeather,
		Leggings:  armour.TierLeather,
		Boots:     armour.TierLeather,
	}
}

func (*Archer) Handler(p *Player) player.Handler {
	return &ArcherHandler{p: p}
}

type ArcherHandler struct {
	player.NopHandler
	p *Player
}

func (*Archer) Tickers(p *Player) []*TickerFunc {
	return nil
}

func (handler *ArcherHandler) HandleItemUse(ctx *event.Context) {
	player := handler.p
	if archer, ok := player.Class().(*Archer); ok {
		heldItem, _ := player.HeldItems()
		if i, ok := archerItems.Load(heldItem.Item()); ok {
			i := i.(ClassUseItem)
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
