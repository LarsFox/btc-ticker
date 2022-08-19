package streamer

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/LarsFox/btc-ticker/storage"
	"github.com/LarsFox/btc-ticker/ticker"
	"github.com/LarsFox/btc-ticker/ticker/mocks"
)

func TestStreamer(t *testing.T) {
	s := New(func(f []float64) float64 { return 100 }, storage.New(), time.Second/10)

	ctrl := gomock.NewController(t)
	mock := mocks.NewMockPriceStreamSubscriber(ctrl)

	pricesCh := make(chan ticker.TickerPrice)
	errCh := make(chan error)

	mock.EXPECT().SubscribePriceStream(gomock.Any()).Return(
		pricesCh, errCh,
	)

	s.Subscribe(mock)
	go s.ProduceFairPrice()

	// Stopping the test early.
	go func() {
		time.Sleep(time.Second * 6)
		close(s.pricesCh)
	}()

	time.Sleep(time.Second / 2)
	pricesCh <- ticker.TickerPrice{}

	time.Sleep(time.Second / 2)
	pricesCh <- ticker.TickerPrice{Ticker: ticker.BTCUSDTicker}

	time.Sleep(time.Second / 2)
	pricesCh <- ticker.TickerPrice{Ticker: ticker.BTCUSDTicker, Price: "0"}

	time.Sleep(time.Second / 2)
	pricesCh <- ticker.TickerPrice{Ticker: ticker.BTCUSDTicker, Price: "100.123", Time: time.Now()}

	time.Sleep(time.Second / 2)
	errCh <- errors.New("the greatest error that make tickers tremble")

	s.Listen()
}
