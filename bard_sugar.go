package hcf

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type BardUseSugar struct{}

func (BardUseSugar) Energy() int {
	return 45
}
func (BardUseSugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 3, 10*time.Second)
}
func (BardUseSugar) AffectMates() bool   { return true }
func (BardUseSugar) AffectEnemies() bool { return false }
func (BardUseSugar) Item() world.Item    { return item.Sugar{} }

type BardHeldSugar struct{}

func (BardHeldSugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 2, 5*time.Second)
}
func (BardHeldSugar) Item() world.Item { return item.Sugar{} }
