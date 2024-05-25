package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string // endpoint
		val []byte // content
	}{
		{
			key: "https://example.com",
			val: []byte("test-data"),
		},
		{
			key: "",
			val: []byte(""),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)

			err := cache.Add(c.key, c.val)
			if c.key == "" {
				if err == nil {
					t.Errorf("expected error caused by empty key added")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error adding key %v in test case %v: %v", c.key, i, err)
				return
			}

			got, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected key %v to be found in test case %v", c.key, i)
				return
			}

			if string(got) != string(c.val) {
				t.Errorf("expected value %v, but got %v in test %v", c.val, got, i)
			}
		})

	}
}
