package pokecache

import (
	"fmt"
	"testing"
	"time"
)

type testCase struct {
	key string
	val []byte
}

type testFunc = func(cases []testCase, cache *Cache[[]byte])

func TestAdd(t *testing.T) {
	testCache(func(cases []testCase, cache *Cache[[]byte]) {
		for i, c := range cases {
			t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
				cache.Add(c.key, &c.val)
				_, ok := cache.Entries[c.key]
				if !ok {
					t.Errorf("Expected to find key: %s", c.key)
					return
				}
				fmt.Println("PASS")
			})
		}
	})
}

func TestReapLoop(t *testing.T) {
	testCache(func(cases []testCase, cache *Cache[[]byte]) {
		const waitTime = time.Millisecond * 10
		testCase := cases[0]
		cache.Add(testCase.key, &testCase.val)

		_, ok := cache.Get(testCase.key)
		if !ok {
			t.Errorf("Expecting to find key: %s", testCase.key)
			return
		}

		time.Sleep(waitTime)

		_, ok = cache.Get(testCase.key)
		if ok {
			t.Errorf("Expecting not to find key: %s", testCase.key)
			return
		}

		fmt.Println("PASS")
	})
}

func TestGet(t *testing.T) {
	testCache(func(cases []testCase, cache *Cache[[]byte]) {
		for i, c := range cases {
			cache.Entries[c.key] = cachEntry[[]byte]{
				createdAt: nil,
				val:       &c.val,
			}
			t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
				_, found := cache.Get(c.key)
				if !found {
					t.Errorf("Expected to find key: %s", c.key)
					return
				}
				fmt.Println("PASS")
			})
		}
		t.Run(fmt.Sprintf("Test case %v", len(cases)+1), func(t *testing.T) {
			missingKey := "missingKey"
			_, found := cache.Get(missingKey)
			if found {
				t.Errorf("Not expected to find key: %s", missingKey)
				return
			}
			fmt.Println("PASS")
		})
	})
}

func testCache(tf testFunc) {
	interval := time.Millisecond * 5
	cache := NewCache[[]byte](interval, true)
	cases := []testCase{
		{
			key: "https://example.com",
			val: []byte("some data"),
		},
		{
			key: "https://example.com/next_page",
			val: []byte("some more data"),
		},
	}
	tf(cases, cache)
}
