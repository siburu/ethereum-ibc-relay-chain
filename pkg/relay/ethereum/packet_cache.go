package ethereum

import math "math"

type packetCache[Event any] map[uint64]packetCacheEntry[Event]

type packetCacheEntry[Event any] struct {
	height uint64
	event  Event
}

func newPacketCache[Event any]() packetCache[Event] {
	return make(map[uint64]packetCacheEntry[Event])
}

func (c packetCache[Event]) add(seq uint64, height uint64, event Event) {
	if _, ok := c[seq]; ok {
		panic("an entry already exists")
	}
	c[seq] = packetCacheEntry[Event]{height, event}
}

func (c packetCache[Event]) evict(seqs []uint64) {
	for _, seq := range seqs {
		delete(c, seq)
	}
}

func (c packetCache[Event]) calcMinimumHeight() uint64 {
	minHeight := uint64(math.MaxUint64)
	for _, e := range c {
		if e.height < minHeight {
			minHeight = e.height
		}
	}
	return minHeight
}
