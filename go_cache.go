package go_cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[any]cacheItem
	mutex sync.RWMutex
}
type cacheItem struct {
	value      any
	expiration time.Time
}

const defaultValue = "default value"

func NewCache() *Cache {
	return &Cache{items: make(map[any]cacheItem)}
}
func (c *Cache) Get(key any) (any, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.items[key]
	if ok {
		if time.Now().After(val.expiration) {
			delete(c.items, key)
			return nil, errors.New("key is expired")
		}
		return val.value, nil
	}
	return defaultValue, errors.New("key not found")
}
func (c *Cache) Set(key, value any, duration time.Duration) {
	expiration := time.Now().Add(duration)
	item := cacheItem{
		value:      value,
		expiration: expiration,
	}
	c.mutex.Lock()
	c.items[key] = item
	c.mutex.Unlock()
}
func (c *Cache) Remove(key any) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.items[key]
	if ok {
		delete(c.items, key)
		return true, nil
	}
	return false, errors.New("key not found")
}
