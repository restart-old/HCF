package hcf

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
)

type ArmourTiers struct {
	Helmet    armour.Tier
	Chestlate armour.Tier
	Leggings  armour.Tier
	Boots     armour.Tier
}

func IsClass(a item.ArmourContainer, class Class) bool {
	tiers := class.ArmourTiers()

	helmet := item.Helmet{Tier: tiers.Helmet}
	chestplate := item.Chestplate{Tier: tiers.Chestlate}
	leggings := item.Leggings{Tier: tiers.Leggings}
	boots := item.Boots{Tier: tiers.Boots}
	return a.Helmet().Item() == helmet && a.Chestplate().Item() == chestplate && a.Leggings().Item() == leggings && a.Boots().Item() == boots
}
