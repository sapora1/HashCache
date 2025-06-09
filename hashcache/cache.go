package hashcache

import (
	"container/list"
	"context"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

type Cache struct {
	capacity    int
	store       map[string]*list.Element
	list        *list.List
	mutex       sync.Mutex
	redisClient *redis.Client
	useRedis    bool
}

type entry struct {
	key   string
	value interface{}
}

func NewCreateCache(capacity int, redisOptions *redis.Options) *Cache {

	var client *redis.Client
	useRedis := false
	if redisOptions != nil {
		client = redis.NewClient(redisOptions)
		useRedis = true
	}
	return &Cache{
		capacity:    capacity,
		store:       make(map[string]*list.Element),
		list:        list.New(),
		redisClient: client,
		useRedis:    useRedis,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, exist := c.store[key]; exist {
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}

	if c.useRedis {
		val, err := c.redisClient.Get(context.Background(), key).Result()
		if err == nil {
			// Insert into local cache
			c.Put(key, val)
			return val, true
		}
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

	if c.useRedis {
		c.redisClient.Set(context.Background(), key, value, 0)
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, exist := c.store[key]; exist {
		c.list.Remove(elem)
		delete(c.store, key)
	}
}

func (c *Cache) Len() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.store)
}
