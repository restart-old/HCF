package bard

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type UseIronIngot struct{}

func (UseIronIngot) Energy() int {
	return 45
}
func (UseIronIngot) Effect() effect.Effect {
	return effect.New(effect.Resistance{}, 3, 10*time.Second)
}
func (UseIronIngot) AffectMates() bool   { return true }
func (UseIronIngot) AffectEnemies() bool { return false }
func (UseIronIngot) Item() world.Item    { return item.IronIngot{} }

type HeldIronIngot struct{}

func (HeldIronIngot) Effect() effect.Effect {
	return effect.New(effect.Resistance{}, 2, 5*time.Second)
}
func (HeldIronIngot) Item() world.Item { return item.IronIngot{} }
