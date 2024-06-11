package pokecache

import "time"

func (c *Cache) reapLoop() {
	for {
		c.mut.Lock()
		for k, v := range c.entries {
			if time.Since(v.createdAt) > c.interval {
				delete(c.entries, k)
			}
		}
		c.mut.Unlock()
	}
}
