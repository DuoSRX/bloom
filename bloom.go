package bloom

import (
	"hash/fnv"
)

type BloomFilter struct {
	values   map[uint64]bool
	capacity uint64
}

func str_to_fnv(s string) (uint64, uint64) {
	h1 := fnv.New64()
	h1.Write([]byte(s))

	h2 := fnv.New64a()
	h2.Write([]byte(s))

	return h1.Sum64(), h2.Sum64()
}

func (bf *BloomFilter) set(h uint64, value bool) {
	bf.values[h%bf.capacity] = value
}

func (bf *BloomFilter) get(h uint64) bool {
	return bf.values[h%bf.capacity]
}

func New(capacity int) *BloomFilter {
	return &BloomFilter{make(map[uint64]bool, capacity), uint64(capacity)}
}

func (bf *BloomFilter) Add(s string) {
	h1, h2 := str_to_fnv(s)
	bf.set(h1, true)
	bf.set(h2, true)
}

func (bf *BloomFilter) IsPresent(s string) bool {
	h1, h2 := str_to_fnv(s)

	return bf.get(h1) && bf.get(h2)
}
