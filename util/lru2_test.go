package util_test

import (
	"github.com/johnny4fun/TermDict/util"
	"strconv"
	"strings"
	"testing"
)

func TestLRU(t *testing.T) {
	t.Run("empty LRU pop", func(t *testing.T) {
		lru := util.NewLRUCache2(10)

		obj := lru.PopFront()

		assertObject(t, obj, nil)
	})

	t.Run("LRU add", func(t *testing.T) {
		lru := util.NewLRUCache2(10)

		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertLRUSize(t, lru, 10)

		want := strings.Split("0,1,2,3,4,5,6,7,8,9",",")
		for i := 0; i < 10; i++ {
			assertLRUContent(t, lru, want)
		}

		t.Run("LRU add more the capacity limit", func(t *testing.T) {
			for i := 10; i < 20; i++ {
				obj := &TestObject{strconv.Itoa(i)}
				lru.Add(obj)
			}

			assertLRUSize(t, lru, 10)
			for i := 0; i < 10; i++ {
				assertLRUNotContains(t, lru, strconv.Itoa(i))
			}
			for i := 10; i < 20; i++ {
				assertLRUContains(t, lru, strconv.Itoa(i))
			}
		})
	})

	t.Run("query when lru is not empty", func(t *testing.T) {
		lru := util.NewLRUCache2(10)
		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertLRUSize(t, lru, 10)

		// query the oldest one
		lru.Get("0")
		lru.Get("1")
		var want []string
		want = strings.Split("2,3,4,5,6,7,8,9,0,1", ",")
		assertLRUContent(t, lru, want)

		// query item in mid
		lru.Get("9")
		want = strings.Split("2,3,4,5,6,7,8,0,1,9", ",")
		assertLRUContent(t, lru, want)

		// query last item
		lru.Get("9")

		want = strings.Split("2,3,4,5,6,7,8,0,1,9",",")
		assertLRUContent(t, lru, want)
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

func assertLRUContent(t *testing.T, lru *util.LRUCache2, want []string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	if got, want := len(objs), len(want); got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	for i := uint(0); i < lru.Size(); i++ {
		if got, want := objs[i].Key(), want[i]; got != want {
			t.Errorf("item[%d]: got %q, want %q", i, got, want)
		}
	}

}

func assertObject(t *testing.T, obj util.Object, expect util.Object) {
	t.Helper()
	if obj != expect {
		t.Errorf("got %v, expected %v", obj, expect)
	}
}

func assertLRUContains(t *testing.T, lru util.LRU, want string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	for _, obj := range objs {
		if obj.Key() == want {
			return
		}
	}

	t.Errorf("%q not expected in %v", want, objs)
}

func assertLRUNotContains(t *testing.T, lru util.LRU, want string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	for _, obj := range objs {
		if obj.Key() == want {
			t.Fatalf("%q not expected in %v", want, objs)
		}
	}
}

func assertLRUSize(t *testing.T, lru util.LRU, want uint) {
	t.Helper()
	if got := lru.Size(); got != want {
		t.Errorf("LRU size is %d, expect %d", got, want)
	}
}
