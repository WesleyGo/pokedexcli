package pokecache

func (c *Cache) GetFromCache(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	entry, ok := c.entries[key]

	return entry.val, ok
}
