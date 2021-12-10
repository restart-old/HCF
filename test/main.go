package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/dragonfly-on-steroids/hcf"
)

func main() {
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	session := hcf.NewSession()
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}

		P := hcf.NewPlayer(p)
		p.Armour().(*inventory.Armour).Inv().Handle(&hcf.ClassHandler{P: P})
		p.Inventory().AddItem(item.NewStack(item.Sugar{}, 64))
		p.Inventory().AddItem(item.NewStack(item.SpiderEye{}, 64))

		p.Inventory().AddItem(item.NewStack(item.Helmet{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Chestplate{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Leggings{Tier: armour.TierGold}, 1))
		p.Inventory().AddItem(item.NewStack(item.Boots{Tier: armour.TierGold}, 1))
		session.StorePlayer(P)
	}
}
