package cache_test

import (
	"github.com/Johnny4Fun/TermDict/util/cache"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestLRUThreadSafe(t *testing.T) {
	t.Run("Cache add", func(t *testing.T) {
		lru := cache.NewThreadSafeLRU(cache.NewLRUCache(10))

		batchAddMultiThread(lru, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

		within(t, 5*time.Second, func() {
			checkLoop(5*time.Millisecond, func() bool {
				return lru.Size() == 10
			})
		})

		for i := 0; i < 10; i++ {
			assertLRUContains(t, lru, strconv.Itoa(i))
		}

		t.Run("add when full", func(t *testing.T) {

			wg := batchAddMultiThread(lru, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)

			within(t, 5*time.Second, func() {
				wg.Wait()
			})

			for i := 10; i < 20; i++ {
				assertLRUContains(t, lru, strconv.Itoa(i))
			}

		})
	})

	t.Run("empty Cache pop", func(t *testing.T) {
		lru := cache.NewThreadSafeLRU(cache.NewLRUCache(10))
		doConcurrently(func(index int) {
			obj := lru.PopFront()
			assertObject(t, obj, nil)
		}, 0, 1, 2, 3)
	})

	t.Run("query when cache is either not empty or not full", func(t *testing.T) {
		lru := cache.NewLRUCache(20)
		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertCacheSize(t, lru, 10)

		t.Run("query item from start/mid/tail", func(t *testing.T) {
			within(t, 5*time.Second, func() {
				wg := doConcurrently(func(index int) {
					lru.Get(strconv.Itoa(index))
				}, 0, 5, 9)
				wg.Wait()
			})

			want := strings.Split("0,5,9", ",")
			assertCacheSize(t, lru, 10)
			assertLRUTailContent(t, lru, want)

		})

	})
	t.Run("query when cache is full", func(t *testing.T) {
		lru := cache.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertCacheSize(t, lru, 10)

		t.Run("query item from start/mid/tail", func(t *testing.T) {
			within(t, 5*time.Second, func() {
				wg := doConcurrently(func(index int) {
					lru.Get(strconv.Itoa(index))
				}, 0, 5, 9)
				wg.Wait()
			})

			want := strings.Split("0,5,9", ",")
			assertLRUTailContent(t, lru, want)

		})

	})
}
