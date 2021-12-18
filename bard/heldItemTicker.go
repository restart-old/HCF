package bard

import (
	"github.com/dragonfly-on-steroids/hcf"
)

func heldItemFunc(b *Bard, p *hcf.Player) func() {
	return func() {
		heldItem, _ := p.HeldItems()
		if i, ok := bardHeldItems.Load(heldItem.Item()); ok {
			i := i.(hcf.ClassHeldItem)
			for _, e := range p.PlayersInRadius(b.effectRadius) {
				e.AddEffect(i.Effect())
			}
		}
	}
}
