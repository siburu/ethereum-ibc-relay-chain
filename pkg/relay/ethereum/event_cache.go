package ethereum

import (
	math "math"
	"sort"
)

type eventCache[Event any] struct {
	entryOfSeq      map[uint64]*eventCacheEntry[Event]
	entriesOfHeight map[uint64][]*eventCacheEntry[Event]
}

type eventCacheEntry[Event any] struct {
	height   uint64
	sequence uint64
	event    Event
}

func newEventCache[Event any]() eventCache[Event] {
	return eventCache[Event]{
		entryOfSeq:      make(map[uint64]*eventCacheEntry[Event]),
		entriesOfHeight: make(map[uint64][]*eventCacheEntry[Event]),
	}
}

func (c eventCache[Event]) add(height uint64, seq uint64, event Event) {
	if _, ok := c.entryOfSeq[seq]; ok {
		return
	}
	entry := &eventCacheEntry[Event]{height, seq, event}
	c.entryOfSeq[seq] = entry
	c.entriesOfHeight[height] = append(c.entriesOfHeight[height], entry)
}

func (c eventCache[Event]) remove(seq uint64) {
	entry, ok := c.entryOfSeq[seq]
	if !ok {
		return
	}
	delete(c.entryOfSeq, seq)
	entries := c.entriesOfHeight[entry.height]
	for i, e := range entries {
		if e == entry {
			entries = append(entries[:i], entries[i+1:]...)
			if len(entries) == 0 {
				delete(c.entriesOfHeight, entry.height)
			} else {
				c.entriesOfHeight[entry.height] = entries
			}
			return
		}
	}
	panic("entry not found in entriesOfHeight")

}

func (c eventCache[Event]) getEvent(seq uint64) (Event, bool) {
	if entry, ok := c.entryOfSeq[seq]; ok {
		return entry.event, true
	} else {
		var zero Event
		return zero, false
	}
}

type heights []uint64

func (a heights) Len() int           { return len(a) }
func (a heights) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a heights) Less(i, j int) bool { return a[i] < a[j] }

// getAllEvents returns all events ordered by height
func (c eventCache[Event]) getAllEvents() []Event {
	// get sorted heights
	var heights heights
	for h, _ := range c.entriesOfHeight {
		heights = append(heights, h)
	}
	sort.Sort(heights)

	// get all events ordered by height
	var events []Event
	for _, h := range heights {
		for _, entry := range c.entriesOfHeight[h] {
			events = append(events, entry.event)
		}
	}
	return events
}

func (c eventCache[Event]) calcMinimumHeight() uint64 {
	minHeight := uint64(math.MaxUint64)
	for h, _ := range c.entriesOfHeight {
		if h < minHeight {
			minHeight = h
		}
	}
	return minHeight
}
