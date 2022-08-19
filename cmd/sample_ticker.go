package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LarsFox/btc-ticker/ticker"
)

type sampleTicker struct {
	ticker   ticker.Ticker
	pricesCh chan ticker.TickerPrice
	errCh    chan error
}

func newSampleTicker() *sampleTicker {
	s := &sampleTicker{
		ticker:   ticker.BTCUSDTicker,
		pricesCh: make(chan ticker.TickerPrice),
		errCh:    make(chan error),
	}

	return s
}

// tick sends tickers forever and ever.
func (s *sampleTicker) tick(d time.Duration) {
	for t := range time.Tick(d) {
		price := 93 + rand.Float64()*13
		s.pricesCh <- ticker.TickerPrice{
			Ticker: s.ticker,
			Time:   t,
			Price:  fmt.Sprintf("%.4f", price),
		}
	}
}

// tickAndDie ticks n times, then meets its end.
func (s *sampleTicker) tickAndDie(d time.Duration, n int) {
	var i int
	for t := range time.Tick(d) {
		price := 78 + rand.Float64()*44
		s.pricesCh <- ticker.TickerPrice{
			Ticker: s.ticker,
			Time:   t,
			Price:  fmt.Sprintf("%.4f", price),
		}

		if i == n {
			s.errCh <- fmt.Errorf("RIP ticker after %d ticks", n)
			close(s.pricesCh)
			close(s.errCh)
			return
		}

		i++
	}
}

// SubscribePriceStream returns the ticker channels to subscribe.
func (s *sampleTicker) SubscribePriceStream(ticker.Ticker) (chan ticker.TickerPrice, chan error) {
	return s.pricesCh, s.errCh
}
