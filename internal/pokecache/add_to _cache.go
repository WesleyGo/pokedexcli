package pokecache

import "time"

func (c *Cache) AddToCache(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
