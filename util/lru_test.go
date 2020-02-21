package util_test

import (
	"strconv"
	"testing"
)

type TestObject struct {
	key string
}

func (o *TestObject) Key() string {
	return o.key
}

func TestLRU(t *testing.T) {
	t.Run("empty LRU pop", func(t *testing.T) {
		lru := util.NewLRUCache(10)

		obj := lru.PopFront()

		assertObject(t, obj, nil)
	})

	t.Run("LRU add", func(t *testing.T) {
		lru := util.NewLRUCache(10)

		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertLRUSize(t, lru, 10)

		storage := lru.GetStorageMap()
		for i := 0; i < 10; i++ {
			assertMapContains(t, storage, strconv.Itoa(i))
		}

		t.Run("LRU add more the capacity limit", func(t *testing.T) {
			for i := 10; i < 20; i++ {
				obj := &TestObject{strconv.Itoa(i)}
				lru.Add(obj)
			}

			assertLRUSize(t, lru, 10)
			for i := 0; i < 10; i++ {
				assertMapNotContains(t, storage, strconv.Itoa(i))
			}
			for i := 10; i < 20; i++ {
				assertMapContains(t, storage, strconv.Itoa(i))
			}
		})
	})

	t.Run("query when lru is not empty", func(t *testing.T) {
		lru := util.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertLRUSize(t, lru, 10)

		// query the oldest one
		lru.Get("0")
		lru.Get("1")

		// query item in mid
		lru.Get("9")

		// query last item
		lru.Get("9")

		want := "2345678019"
		// pop all
		got := [10]byte{}
		for i := 0; i < 10; i++ {
			obj := lru.PopFront()
			got[i] = obj.Key()[0]
		}

		if string(got[:]) != want {
			t.Errorf("got %v, want %s", got, want)
		}


	})

	t.Run("clear lru cache", func(t *testing.T) {
		lru := util.NewLRUCache(10)
		lru.Add(&TestObject{"nmsl"})

		lru.ClearWithFunc(nil)
		storage := lru.GetStorageMap()
		if l := len(storage); l != 0 {
			t.Errorf("storage is not empty, expected length %d", 0)
		}

		if l := lru.Size(); l != 0 {
			t.Errorf("size is %d, expected length %d", l, 0)
		}
	})
}

func assertObject(t *testing.T, obj util.Object, expect util.Object) {
	t.Helper()
	if obj != expect {
		t.Errorf("got %v, expected %v", obj, expect)
	}
}

func assertMapContains(t *testing.T, storageMap map[string]*util.Node, want string) {
	t.Helper()
	if _, ok := storageMap[want]; !ok {
		t.Errorf("key %q expected in %v", want, storageMap)
	}
}

func assertMapNotContains(t *testing.T, storageMap map[string]*util.Node, want string) {
	t.Helper()
	if _, ok := storageMap[want]; ok {
		t.Errorf("key %q expected NOT in %v", want, storageMap)
	}
}

func assertLRUSize(t *testing.T, lru *util.LRUCache, want uint) {
	t.Helper()
	if got := lru.Size(); got != want {
		t.Errorf("LRU size is %d, expect %d", got, want)
	}
}
