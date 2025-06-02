package hashcache

import (
	"container/list"
	"sync"
)

type Cache struct {
	capacity int
	store    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

type entry struct {
	key   string
	value interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, exist := c.store[key]; exist {
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

func (c *Cache) Put(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, exist := c.store[key]; exist {
		c.list.MoveToFront(elem)
		return
	}

	if c.capacity <= c.list.Len() {
		last := c.list.Back()
		c.list.Remove(last)
		delete(c.store, last.Value.(*entry).key)
	}

	elem := c.list.PushFront(&entry{key, value})
	c.store[key] = elem
}
