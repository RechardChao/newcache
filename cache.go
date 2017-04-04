package newCache

import "sync"

type myCache interface {
	Get(key interface{}) (interface{}, error)
	Del(key interface{})
	Put(key, val interface{})
}

func NewmyCache(newElemFunc func(key interface{}) (interface{}, error)) *newCache {
	c := &newCache{
		mu:          &sync.Mutex{},
		memo:        make(map[interface{}]*newCacheElem),
		newElemFunc: newElemFunc,
	}
	return c
}
