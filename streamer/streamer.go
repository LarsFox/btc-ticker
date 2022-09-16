package streamer

import (
	"log"
	"strconv"
	"time"

	"github.com/LarsFox/btc-ticker/storage"
	"github.com/LarsFox/btc-ticker/ticker"
)

// Streamer collects data from the tickers, calculates the fair price
// and prints it every period.
//
// The fairness of calculated price depends on the passed calc function.
type Streamer struct {
	calc     func([]float64) float64
	period   time.Duration
	pricesCh chan ticker.TickerPrice
	storage  *storage.Storage
}

func New(calc func([]float64) float64, s *storage.Storage, period time.Duration) *Streamer {
	return &Streamer{
		calc:     calc,
		period:   period,
		pricesCh: make(chan ticker.TickerPrice, 100),
		storage:  s,
	}
}

// Subscribe handles the messages from a single ticker.PriceStreamSubscriber.
func (s *Streamer) Subscribe(p ticker.PriceStreamSubscriber) {
	ch, errCh := p.SubscribePriceStream(ticker.BTCUSDTicker)

	go func() {
		for {
			select {
			case price := <-ch:
				s.pricesCh <- price
			case err := <-errCh:
				log.Println("PriceStreamSubscriber err:", err)
				return
			}
		}
	}()
}

// Listen consumes all subscribers messages.
func (s *Streamer) Listen() {
	for tickerPrice := range s.pricesCh {
		if tickerPrice.Ticker != ticker.BTCUSDTicker {
			continue
		}

		f, err := strconv.ParseFloat(tickerPrice.Price, 64)
		if err != nil {
			log.Printf("invalid price from ticker: %s", tickerPrice.Price)
			continue
		}

		// Most likely an error.
		if f == 0 {
			continue
		}

		s.storage.Add(tickerPrice.Time, f)
	}
}

// DelayMinute is a helper function to start at the exact minute start.
func (s *Streamer) DelayMinute() {
	for {
		now := time.Now()
		if now.Second() == 0 {
			return
		}

		time.Sleep(time.Second)
	}
}

// ProduceFairPrice prints calculated price for the previous period.
// First print is skipped due to the time.Tick initial delay.
// This skip also proves useful as there is not enough data for the first print.
func (s *Streamer) ProduceFairPrice() {
	for t := range time.Tick(s.period) {
		minute := t.Add(-time.Minute)

		prices := s.storage.Retrieve(minute)

		val := s.calc(prices)

		log.Println(minute.Unix(), val)
	}
}
