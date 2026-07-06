package lrucache

import (
	"math/rand"
	"testing"
)

func BenchmarkLRUCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := 1
		r := rand.New(rand.NewSource(42))
		cache := NewLRUCache[int, int](40)
		for j := 0; j < 10000; j++ {
			ri := r.Intn(10)
			if ri == 0 {
				cache.Put(k, k)
				k++
			} else {
				cache.Get(r.Intn(k))
			}
		}
	}
}
