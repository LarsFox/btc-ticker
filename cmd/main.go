package main

import (
	"math/rand"
	"time"

	"github.com/LarsFox/btc-ticker/calc"
	"github.com/LarsFox/btc-ticker/storage"
	"github.com/LarsFox/btc-ticker/streamer"
	"github.com/LarsFox/btc-ticker/ticker"
)

const (
	tickersNum        = 100
	healthyTickersNum = tickersNum * 0.8
)

func mockSampleTickers() []ticker.PriceStreamSubscriber {
	result := make([]ticker.PriceStreamSubscriber, 0, tickersNum)

	for i := 0; i < tickersNum; i++ {
		s := newSampleTicker()
		result = append(result, s)

		if i < healthyTickersNum {
			go s.tick(time.Second * 2)
			continue
		}

		go s.tickAndDie(time.Second*2, i)
	}

	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tickers := mockSampleTickers()

	s := streamer.New(
		calc.Median,
		storage.New(),
		time.Minute,
	)

	for _, ticker := range tickers {
		s.Subscribe(ticker)
	}

	go s.Listen()

	s.DelayMinute()
	s.ProduceFairPrice()
}
