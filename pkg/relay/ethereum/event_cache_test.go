package ethereum

import (
	math "math"
	"testing"
)

func TestEventCache(t *testing.T) {
	c := newEventCache[string]()
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}

	// removing inexistent entries is ok
	c.remove(1)
	c.remove(2)
	c.remove(3)
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}

	// add entries
	c.add(1, 1, "a")
	c.add(1, 2, "b")
	c.add(2, 4, "c")
	c.add(3, 3, "d")
	c.add(3, 5, "e")
	if minHeight := c.calcMinimumHeight(); minHeight != 1 {
		t.Fatal(minHeight)
	}
	if event, ok := c.getEvent(3); !ok || event != "d" {
		t.Fatal(event)
	}
	events := c.getAllEvents()
	if nevents := len(events); nevents != 5 {
		t.Fatal(nevents)
	}
	if event := events[0]; event != "a" {
		t.Fatal(event)
	}
	if event := events[1]; event != "b" {
		t.Fatal(event)
	}
	if event := events[2]; event != "c" {
		t.Fatal(event)
	}
	if event := events[3]; event != "d" {
		t.Fatal(event)
	}
	if event := events[4]; event != "e" {
		t.Fatal(event)
	}

	// remove some
	c.remove(1)
	c.remove(5)
	if minHeight := c.calcMinimumHeight(); minHeight != 1 {
		t.Fatal(minHeight)
	}

	// remove all from height=1
	c.remove(2)
	if minHeight := c.calcMinimumHeight(); minHeight != 2 {
		t.Fatal(minHeight)
	}

	// remove all from height=1,2
	c.remove(4)
	if minHeight := c.calcMinimumHeight(); minHeight != 3 {
		t.Fatal(minHeight)
	}

	// remove all entries
	c.remove(3)
	if minHeight := c.calcMinimumHeight(); minHeight != math.MaxUint64 {
		t.Fatal(minHeight)
	}
}
