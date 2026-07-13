package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAllElements(cache *LRUCache[int, int]) ([]int, []int) {
	resKey := make([]int, 0)
	resVal := make([]int, 0)

	for k, v := range cache.cache {
		resKey = append(resKey, k)
		resVal = append(resVal, v.Value.(*entry[int, int]).val)
	}

	return resKey, resVal
}

func TestLRUCache_Put(t *testing.T) {
	cache := NewLRUCache[int, int](20)
	for i := 0; i < 6; i++ {
		cache.Put(i, i)
	}

	k, v := getAllElements(cache)

	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5}, k)
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5}, v)
}

func TestLRUCache_Get(t *testing.T) {
	cache := NewLRUCache[int, int](10)
	for i := 0; i < 8; i++ {
		cache.Put(i, i)
	}
	e, ok := cache.Get(3)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, e)

	for i := 0; i < 8; i++ {
		cache.Put(i+8, i+8)
	}

	e, ok = cache.Get(3)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, e)

	e, ok = cache.Get(2)
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, e)
}

func TestLRUCache_Delete(t *testing.T) {
	cache := NewLRUCache[int, int](10)
	for i := 0; i < 8; i++ {
		cache.Put(i, i)
	}
	e, ok := cache.Get(3)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, e)

	cache.Delete(3)
	e, ok = cache.Get(3)
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, e)
}
