package hashcache

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Init() {
	var err error
	cacheSize, err := strconv.Atoi(os.Getenv("CACHE_SIZE"))
	if err != nil {
		os.Setenv("CACHE_SIZE", "10")
		cacheSize = 20
	}
	log.Println("Cache size is: ", cacheSize)
}

func NewCreateCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		store:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func main() {
	fmt.Print("HashCache")
	Init()
	cache := NewCreateCache(3)
	cache.Put("A", "10")
	cache.Put("B", "11")
	cache.Put("C", "12")
	cache.Put("A", "10")
	cache.Put("D", "13")
	cache.Put("A", "10")
	for key, value := range cache.store {
		fmt.Println(key, value.Value.(*entry).value)
	}
}
