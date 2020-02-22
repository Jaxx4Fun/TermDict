package lru_test

import (
	"github.com/Johnny4Fun/TermDict/util/lru"
	"strconv"
	"strings"
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
		lru := lru.NewLRUCache(10)

		obj := lru.PopFront()

		assertObject(t, obj, nil)
	})

	t.Run("LRU add", func(t *testing.T) {
		lru := lru.NewLRUCache(10)

		for i := 0; i < 10; i++ {
			obj := &TestObject{strconv.Itoa(i)}
			lru.Add(obj)
		}

		assertLRUSize(t, lru, 10)

		want := strings.Split("0,1,2,3,4,5,6,7,8,9", ",")
		assertLRUContent(t, lru, want)

		t.Run("LRU add more the capacity cap", func(t *testing.T) {
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

	//t.Run("query when lru is not empty", func(t *testing.T) {
	lru := lru.NewLRUCache(10)
	for i := 0; i < 10; i++ {
		obj := &TestObject{strconv.Itoa(i)}
		lru.Add(obj)
	}

	assertLRUSize(t, lru, 10)

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
	//})
	//
	//t.Run("clear lru cache", func(t *testing.T) {
	//	lru := lru.NewLRUCache(10)
	//	lru.Add(&TestObject{"nmsl"})
	//
	//	lru.ClearWithFunc(nil)
	//	want := []string{}
	//	assertLRUContent(t, lru, want)
	//})
}

func assertLRUContent(t *testing.T, lru *lru.LRUCache, want []string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	if got, want := len(objs), len(want); got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	for i, obj := range objs {
		if want, got := want[i], obj.Key(); got != want {
			t.Errorf("item[%d]: got %q, want %q", i, got, want)
		}
	}
}

func assertObject(t *testing.T, obj lru.Object, expect lru.Object) {
	t.Helper()
	if obj != expect {
		t.Errorf("got %v, expected %v", obj, expect)
	}
}

func assertLRUContains(t *testing.T, lru lru.LRU, want string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	for _, obj := range objs {
		if obj.Key() == want {
			return
		}
	}

	t.Errorf("%q not expected in %v", want, objs)
}

func assertLRUNotContains(t *testing.T, lru lru.LRU, want string) {
	t.Helper()

	objs := lru.GetContentByOrder()

	for _, obj := range objs {
		if obj.Key() == want {
			t.Fatalf("%q not expected in %v", want, objs)
		}
	}
}

func assertLRUSize(t *testing.T, lru lru.LRU, want uint) {
	t.Helper()
	if got := lru.Size(); got != want {
		t.Errorf("LRU len is %d, expect %d", got, want)
	}
}
