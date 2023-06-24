package go_cache

import (
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	cache := NewCache()
	cache.Set(1, "test", time.Second)
	val, err := cache.Get(1)
	if err != nil {
		t.Errorf("Failed Get() %v", err)
	}
	if val != "test" {
		t.Errorf("Expected %s Got %s", "test", val)
	}

}
func TestCache_Set(t *testing.T) {
	cache := NewCache()
	cache.Set(1, "test", time.Second)
	val, err := cache.Get(1)
	if err != nil {
		t.Errorf("Did not add item to cache")
	}
	if val != "test" {
		t.Errorf("Did not store correct value")
	}
}
func TestCache_Remove(t *testing.T) {
	cache := NewCache()
	cache.Set(1, "test", time.Second)
	ok, err := cache.Remove(1)
	if !ok {
		t.Errorf("Did not remove item from the cache %v", err)
	}
}
