package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type UseSugar struct{}

func (UseSugar) Energy() int {
	return 45
}
func (UseSugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 3, 10*time.Second)
}
func (UseSugar) AffectMates() bool   { return true }
func (UseSugar) AffectEnemies() bool { return false }
func (UseSugar) Item() world.Item    { return item.Sugar{} }

type HeldSugar struct{}

func (HeldSugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 2, 5*time.Second)
}
func (HeldSugar) Item() world.Item { return item.Sugar{} }
