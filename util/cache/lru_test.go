package cache

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

type TestObject int

// type TestObject struct {
// 	key string
// }

// func (o *TestObject) Key() string {
// 	return o.key
// }
// func (o *TestObject) String() string {
// 	return o.Key()
// }

func checkLoop(interval time.Duration, until func() bool) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if until() {
				return
			}
		}
	}
}

func batchAddMultiThread(lru *ThreadSafeLRUWrapper, ins ...int) *sync.WaitGroup {
	return doConcurrently(func(index int) {
		lru.Add(index, index)
	}, ins...)
}

// doConcurrently 包装了创建协程的方式，并传入协程id
func doConcurrently(f func(index int), ids ...int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go func(index int) {
			f(index)
			defer wg.Done()
		}(id)
	}
	return wg
}
func serializeLruContent(lru LRU) (ans []string) {
	var t interface{}
	t = lru
	var underlyingList *list.List
	switch t := t.(type) {
	case *LRUCache:
		underlyingList = t.timeList
	case *ThreadSafeLRUWrapper:
		underlyingList = t.lru.(*LRUCache).timeList
	default:
		panic("lru is not LRU")
	}

	for iter := underlyingList.Front(); iter != nil; iter = iter.Next() {
		o := iter.Value.(entry).obj
		switch t := o.(type) {
		case int:
			ans = append(ans, strconv.Itoa(t))
		default:
			ans = append(ans, fmt.Sprintf("%v", t))
		}
	}
	return
}
func assertLRUTailContent(t *testing.T, lru LRU, want []string) {
	t.Helper()

	objs := serializeLruContent(lru)

	if got, want := len(objs), len(want); got <= want {
		t.Errorf("got %d, want %d", got, want)
	}

	tailObjs := objs[len(objs)-len(want) : len(objs)]
	for _, s := range want {
		if !contains(tailObjs, s) {
			t.Errorf("got %v, want %v", objs, want)
			break
		}
	}

}

func within(t *testing.T, duration time.Duration, f func()) {
	done := make(chan bool, 1)
	go func() {
		f()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(duration):
		t.Errorf("timed out")
	}
}
func TestLRU(t *testing.T) {
	t.Run("empty Cache pop", func(t *testing.T) {
		lru := NewLRUCache(10)

		obj := lru.PopOldest()

		assertObject(t, obj, nil)
	})

	t.Run("Cache add", func(t *testing.T) {
		lru := NewLRUCache(10)

		for i := 0; i < 10; i++ {
			lru.Add(i, i)
		}

		assertCacheSize(t, lru, 10)

		want := strings.Split("0,1,2,3,4,5,6,7,8,9", ",")
		assertLRUContent(t, lru, want)

		t.Run("Cache add more the capacity cap", func(t *testing.T) {
			for i := 10; i < 20; i++ {
				lru.Add(i, i)
			}

			assertCacheSize(t, lru, 10)
			for i := 0; i < 10; i++ {
				assertLRUNotContains(t, lru, strconv.Itoa(i))
			}
			for i := 10; i < 20; i++ {
				assertLRUContains(t, lru, strconv.Itoa(i))
			}
		})
	})

	//t.Run("query when cache is not empty", func(t *testing.T) {
	lru := NewLRUCache(10)
	for i := 0; i < 10; i++ {
		lru.Add(i, i)
	}

	assertCacheSize(t, lru, 10)

	//t.Run("query the oldest one", func(t *testing.T) {
	lru.Get("0")
	lru.Get("1")
	want2 := strings.Split("2,3,4,5,6,7,8,9,0,1", ",")
	assertLRUContent(t, lru, want2)
	//})

	t.Run("query item in mid", func(t *testing.T) {
		lru.Get("9")
		want := strings.Split("2,3,4,5,6,7,8,0,1,9", ",")
		assertLRUContent(t, lru, want)
	})
	t.Run("query last item", func(t *testing.T) {
		lru.Get("9")

		want := strings.Split("2,3,4,5,6,7,8,0,1,9", ",")
		assertLRUContent(t, lru, want)
	})
}

func assertLRUContent(t *testing.T, lru *LRUCache, want []string) {
	t.Helper()

	objs := serializeLruContent(lru)

	for _, wantElem := range want {

		found := false
		for _, cacheElem := range objs {
			if wantElem == cacheElem {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("want %v not found in %v", wantElem, objs)
			break
		}
	}

}

func assertObject(t *testing.T, obj interface{}, expect interface{}) {
	t.Helper()
	if !reflect.DeepEqual(obj, expect) {
		t.Errorf("got %v, expected %v", obj, expect)
	}
}

func assertLRUContains(t *testing.T, lru LRU, want string) {
	t.Helper()

	objs := serializeLruContent(lru)

	if !contains(objs, want) {
		t.Errorf("%q not expected in %v", want, objs)
	}

}

func contains(objs []string, want string) bool {
	for _, obj := range objs {
		if reflect.DeepEqual(obj, want) {
			return true
		}
	}
	return false
}

func assertLRUNotContains(t *testing.T, lru LRU, want string) {
	t.Helper()

	objs := serializeLruContent(lru)

	for _, obj := range objs {
		if obj == want {
			t.Fatalf("%q not expected in %v", want, objs)
		}
	}
}

func assertCacheSize(t *testing.T, lru ICache, want uint) {
	t.Helper()
	if got := lru.Size(); got != want {
		t.Errorf("Cache len is %d, expect %d", got, want)
	}
}
