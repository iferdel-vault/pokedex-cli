package pokecache

import (
        "testing"
        "fmt"
)

func TestAddGet(t *testing.T) {
    cases := []struct {
        key string // endpoint
        val []byte // content
    } {
        {
            key: "https://example.com",
            val: []byte("test-data"),
        },

    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
        cache := NewCache()
        cache.Add(c.key, c.val)

        got, ok := cache.Get(c.key)
        if !ok {
            t.Errorf("expected a key to be found in test case %v", i)
            return
        }

        if got != c.val {
            t.Errorf("expected value %v, but got %v in test %v", c.val, got, i) 
        }
    })

    }
}
