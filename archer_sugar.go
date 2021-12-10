package hcf

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type ArcherSugar struct{}

func (ArcherSugar) Energy() int {
	return 45
}
func (ArcherSugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 4, 10*time.Second)
}
func (ArcherSugar) AffectMates() bool   { return false }
func (ArcherSugar) AffectEnemies() bool { return false }
func (ArcherSugar) Item() world.Item    { return item.Sugar{} }
