package storage

import (
	"sync"
	"sync/atomic"
	"time"
)

type Storage struct {
	mu        sync.RWMutex
	arrays    map[int64][]float64
	threshold int64
}

func New() *Storage {
	return &Storage{
		arrays: map[int64][]float64{},
	}
}

func minuteStart(t time.Time) int64 {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		0,
		0,
		t.Location(),
	).Unix()
}

// Retrieve takes a value from Storage and deletes it,
// while also setting the threshold to this value.
//
// All the values below the threshold are therefore ignored.
func (s *Storage) Retrieve(t time.Time) []float64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	min := minuteStart(t)

	val := s.arrays[min]
	delete(s.arrays, min)

	atomic.StoreInt64(&s.threshold, min)

	return val
}

// Adds adds the value if it is above the threshold.
func (s *Storage) Add(t time.Time, price float64) {
	if t.Unix() < atomic.LoadInt64(&s.threshold) {
		return
	}

	min := minuteStart(t)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.arrays[min] = append(s.arrays[min], price)
}
