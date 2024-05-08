package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	cache map[string]cacheEntry
	mux  *sync.Mutex
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {

	c:= Cache{
		cache:make(map[string]cacheEntry),
		mux:&sync.Mutex{},
	}
	go c.RemoveLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte)  {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:  val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheEntry,found := c.cache[key]

	return cacheEntry.val,found
}

func (c *Cache) RemoveLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Remove(interval)
	}
}

func (c *Cache) Remove(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k,v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache,k)
		}
		
	}
	
}