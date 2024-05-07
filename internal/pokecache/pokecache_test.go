package pokecache

import (
	"testing"
	"time"
)


func TestCreateCache(t *testing.T) {

	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key1",
			inputVal: []byte("val2"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)

		actual, found := cache.Get(c.inputKey)

		if !found {
			t.Errorf("%v not found",c.inputKey)
			continue
		}

		if string(actual) != string(c.inputVal) {
			t.Errorf("%v does not match %v",string(actual),string(c.inputVal))
			continue
		}
	}
}

func TestRemoveCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	
	cache.Add("key1",[]byte("val1"))
	time.Sleep(interval + (time.Millisecond * 5))

	_,ok := cache.Get("key1")

	if ok {
		t.Errorf("%v it should have been deleted","val1")
	}
	
}