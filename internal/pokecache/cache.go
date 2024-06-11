package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mut      *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) *Cache {
	cache := Cache{
		interval: timeout,
		mut:      &sync.Mutex{},
		entries:  make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return &cache
}
