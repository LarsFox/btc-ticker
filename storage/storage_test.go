package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	s := New()

	now := time.Date(2022, 8, 19, 22, 0, 23, 0, time.Now().Location())
	later := now.Add(time.Minute)

	s.Add(now.Add(time.Second), 1)
	s.Add(now.Add(time.Second*5), 2)
	s.Add(now, 3)

	s.Add(later, 7)
	s.Add(later, 8)
	s.Add(later, 9)

	assert.Equal(t, 2, len(s.arrays))

	val := s.Retrieve(later)

	assert.Equal(t, []float64{7, 8, 9}, val)
	assert.Equal(t, 1, len(s.arrays))

	s.Add(now.Add(time.Second), 1)
	s.Add(now.Add(time.Second*5), 2)
	s.Add(now, 3)

	assert.Equal(t, 1, len(s.arrays))
}
