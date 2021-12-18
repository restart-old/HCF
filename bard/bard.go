package bard

import (
	"math"
	"reflect"
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/dragonfly-on-steroids/hcf"

	"github.com/df-mc/dragonfly/server/player"
)

var DefaultBard = &Bard{}

var bardUseItems sync.Map
var bardHeldItems sync.Map

func init() {
	bardUseItems.Store(UseSugar{}.Item(), UseSugar{})
	bardUseItems.Store(SpiderEye{}.Item(), SpiderEye{})
}

func init() {
	bardHeldItems.Store(HeldSugar{}.Item(), HeldSugar{})
}

func (*Bard) New(p *hcf.Player) hcf.Class {
	b := &Bard{energy: 0, maxEnergy: 120, effectRadius: 35}
	b.AddTicker(hcf.NewTickerFunc(50*time.Millisecond, heldItemFunc(b, p)))
	b.AddTicker(hcf.NewTickerFunc(1*time.Second, func() {
		if b.energy < b.maxEnergy+1 {
			b.energy++
		}
	}))
	return b
}

type Bard struct {
	energy, maxEnergy int
	effectRadius      float64
	effectCoolDown    time.Time
	tickers           []*hcf.TickerFunc
}

func (b *Bard) AddTicker(t *hcf.TickerFunc) {
	b.tickers = append(b.tickers, t)
}

func (b *Bard) Energy() int                       { return b.energy }
func (b *Bard) MaxEnergy() int                    { return b.maxEnergy }
func (b *Bard) RemoveEnergy(energy int)           { b.energy -= energy }
func (b *Bard) EffectRadius() float64             { return b.effectRadius }
func (b *Bard) HasEnoughEnergy(energy int) bool   { return b.energy >= energy }
func (b *Bard) OnEffectCoolDown() bool            { return b.effectCoolDown.After(time.Now()) }
func (b *Bard) SetEffectCoolDown(t time.Duration) { b.effectCoolDown = time.Now().Add(t) }

func (*Bard) Effects() []effect.Effect {
	return []effect.Effect{
		effect.New(effect.Speed{}, 2, 43830*time.Minute),
		effect.New(effect.Regeneration{}, 1, 43830*time.Minute),
		effect.New(effect.Resistance{}, 2, 43830*time.Minute),
	}
}
func (b *Bard) Tickers(p *hcf.Player) []*hcf.TickerFunc {
	return b.tickers
}

func (*Bard) ArmourTiers() hcf.ArmourTiers {
	return hcf.ArmourTiers{
		Helmet:    armour.TierGold,
		Chestlate: armour.TierGold,
		Leggings:  armour.TierGold,
		Boots:     armour.TierGold,
	}
}

func (*Bard) Handler(p *hcf.Player) player.Handler {
	return &Handler{p: p}
}

type Handler struct {
	player.NopHandler
	p *hcf.Player
}

func (handler *Handler) HandleItemUse(ctx *event.Context) {
	var n int
	player := handler.p
	if bard, ok := player.Class().(*Bard); ok {
		heldItem, _ := player.HeldItems()
		if i, ok := bardUseItems.Load(heldItem.Item()); ok {
			i := i.(hcf.ClassUseItem)
			if bard.OnEffectCoolDown() {
				player.Messagef("§cYou cannot use this for another %v seconds!", math.Floor(float64(time.Until(bard.effectCoolDown).Seconds()*10))/10)
				return
			}
			if !bard.HasEnoughEnergy(i.Energy()) {
				player.Messagef("§cYou do not have enough energy for this! You need %v energy. but you only have %v", i.Energy(), bard.Energy())
				return
			}
			bard.SetEffectCoolDown(10 * time.Second)
			bard.RemoveEnergy(i.Energy())
			player.Inventory().RemoveItem(item.NewStack(i.Item(), 1))
			effectName := reflect.TypeOf(i.Effect().Type()).Name()

			for _, p := range player.PlayersInRadius(bard.EffectRadius()) {
				if e, ok := hcf.HasEffectUnderLVL(p, i.Effect(), i.Effect().Level()); ok {
					hcf.NewEffectNowLoss(i.Effect(), e).Add(p)
				} else {
					p.AddEffect(i.Effect())
				}
				n++
			}
			player.Messagef("§eYou have given §9%s %s§e to §a%v §eteammates", effectName, intToRoman(i.Effect().Level()), n)
		}
	}
}

//up to 10
func intToRoman(n int) string {
	m := []string{
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
	}
	return m[n-1]
}
