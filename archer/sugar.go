package archer

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type Sugar struct{}

func (Sugar) Energy() int {
	return 45
}
func (Sugar) Effect() effect.Effect {
	return effect.New(effect.Speed{}, 4, 10*time.Second)
}
func (Sugar) AffectMates() bool   { return false }
func (Sugar) AffectEnemies() bool { return false }
func (Sugar) Item() world.Item    { return item.Sugar{} }
