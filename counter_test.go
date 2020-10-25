package counter

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter100(t *testing.T) {
	var total = 100
	c := NewCounter(0, total)
	defer c.Dispose()
	for i := 0; i < total; i++ {
		c.Plus(1)
	}
	c.Flush()
	assert.Equal(t, total, c.Value())
}

func TestCounter1000(t *testing.T) {
	var total = 1000
	c := NewCounter(0, total)
	defer c.Dispose()
	for i := 0; i < total; i++ {
		c.Plus(1)
	}
	c.Flush()
	assert.Equal(t, total, c.Value())
}

func BenchmarkAtomic(b *testing.B) {
	var total = 10000

	for n := 0; n < b.N; n++ {
		var c = int32(0)
		var cb = func() {
			atomic.AddInt32(&c, 1)
		}
		for i := 0; i < total; i++ {
			go cb()
		}
	}
}

func BenchmarkCounter(b *testing.B) {
	var total = 10000

	for n := 0; n < b.N; n++ {
		var c = NewCounter(0, total)
		for i := 0; i < total; i++ {
			c.Plus(1)
		}
		c.Flush()
	}
}
