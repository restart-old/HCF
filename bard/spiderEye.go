package bard

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type SpiderEye struct{}

func (SpiderEye) Energy() int {
	return 45
}
func (SpiderEye) Effect() effect.Effect {
	return effect.New(effect.Wither{}, 1, 10*time.Second)
}
func (SpiderEye) AffectMates() bool   { return false }
func (SpiderEye) AffectEnemies() bool { return true }
func (SpiderEye) Item() world.Item    { return item.SpiderEye{} }
