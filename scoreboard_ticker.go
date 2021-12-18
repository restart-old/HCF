package hcf

import (
	"time"
)

func ScoreboardTicker(p *Player) TickerFunc {
	return *NewTickerFunc(800*time.Millisecond, func() {
		if len(p.scoreboard.Lines()) <= 0 {
			p.RemoveScoreboard()
		} else {
			p.SendScoreboard(p.scoreboard)
		}
	})
}
