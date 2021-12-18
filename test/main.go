package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/dragonfly-on-steroids/hcf"
	"github.com/dragonfly-on-steroids/hcf/bard"
)

func main() {
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	session := hcf.NewSession()

	hcf.RegisterClass(bard.DefaultBard)
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}

		P := hcf.NewPlayer(p)

		ticker := hcf.ScoreboardTicker(P)
		go ticker.Start()

		p.Armour().Inventory().Handle(&hcf.ClassHandler{P: P})
		p.Inventory().AddItem(item.NewStack(item.Sugar{}, 64))
		p.Inventory().AddItem(item.NewStack(item.SpiderEye{}, 64))

		p.Inventory().AddItem(item.NewStack(item.Helmet{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Chestplate{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Leggings{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Boots{Tier: armour.TierGold}, 1))
		session.StorePlayer(P)
	}
}
