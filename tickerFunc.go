package hcf

import (
	"time"
)

type TickerFunc struct {
	*time.Ticker
	f     func()
	close chan bool
}

func NewTickerFunc(t time.Duration, f func()) *TickerFunc {
	return &TickerFunc{Ticker: time.NewTicker(t), f: f, close: make(chan bool)}
}

func (t *TickerFunc) Start() {
	for {
		select {
		case _, running := <-t.close:
			if !running {
				return
			}
		case <-t.C:
			t.f()
		}
	}
}

func (t *TickerFunc) Stop() {
	t.Ticker.Stop()
	close(t.close)
}
