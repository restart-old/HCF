package hcf

func heldItemFunc(b *Bard, p *Player) func() {
	return func() {
		heldItem, _ := p.HeldItems()
		if i, ok := bardHeldItems.Load(heldItem.Item()); ok {
			i := i.(ClassHeldItem)
			for _, e := range p.PlayersInRadius(b.effectRadius) {
				e.AddEffect(i.Effect())
			}
		}
	}
}
