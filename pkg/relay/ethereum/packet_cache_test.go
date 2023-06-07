package ethereum

import (
	math "math"
	"testing"
)

func TestPacketCache(t *testing.T) {
	c := newPacketCache[string]()
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}
	c.evict([]uint64{1, 2, 3})
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}
	c.add(1, 1, "a")
	c.add(2, 1, "b")
	c.add(3, 2, "c")
	c.add(4, 3, "d")
	c.add(5, 3, "e")
	if minHeight := c.calcMinimumHeight(); minHeight != 1 {
		t.Fatal(minHeight)
	}
	c.evict([]uint64{1, 5})
	if minHeight := c.calcMinimumHeight(); minHeight != 1 {
		t.Fatal(minHeight)
	}
	c.evict([]uint64{2})
	if minHeight := c.calcMinimumHeight(); minHeight != 2 {
		t.Fatal(minHeight)
	}
	c.evict([]uint64{3})
	if minHeight := c.calcMinimumHeight(); minHeight != 3 {
		t.Fatal(minHeight)
	}
	c.evict([]uint64{4})
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}
}
