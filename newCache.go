package newCache

import (
	"sync"
)


type newCacheElem struct {
	
	val interface{}
	
	err error
	
	ready chan struct{}
}


type newCache struct {
	// 全缓存锁
	mu *sync.Mutex

	// 全缓存
	memo map[interface{}]*newCacheElem

	// 重构方法
	newElemFunc func(key interface{}) (interface{}, error)
}

// 获取指定key的缓存，如果缓存不存在则调用重构方法尝试重构。
func (this *newCache) Get(key interface{}) (interface{}, error) {
	this.mu.Lock()
	elem, ok := this.memo[key]
	if !ok {
		
		elem = &newCacheElem{ready: make(chan struct{})}
		this.memo[key] = elem
		this.mu.Unlock()

		
		elem.val, elem.err = this.newElemFunc(key)

		close(elem.ready)
	} else {
		
		this.mu.Unlock()
		
		<-elem.ready
	}
	return elem.val, elem.err
}


func (this *newCache) Del(key interface{}) {
	this.mu.Lock()
	delete(this.memo, key)
	this.mu.Unlock()
}


func (this *newCache) Put(key, val interface{}) {
	this.mx.Lock()
	elem := &newCacheElem{ready: make(chan struct{})}
	this.memo[key] = elem
	this.mu.Unlock()
	elem.val = val
	close(elem.ready)
}
