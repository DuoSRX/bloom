package bloom

import (
	"hash/fnv"
)

type BloomFilter struct {
	values   map[uint64]bool
	capacity uint64
}

func str_to_fnv(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func str_to_fnva(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func New(capacity int) *BloomFilter {
	return &BloomFilter{make(map[uint64]bool, capacity), uint64(capacity)}
}

func (bf *BloomFilter) Add(s string) {
	key := str_to_fnv(s)
	bf.values[key%bf.capacity] = true

	key = str_to_fnva(s)
	bf.values[key%bf.capacity] = true
}

func (bf *BloomFilter) IsPresent(s string) bool {
	h1 := str_to_fnv(s)
	h2 := str_to_fnva(s)

	return bf.values[h1%bf.capacity] && bf.values[h2%bf.capacity]
}
