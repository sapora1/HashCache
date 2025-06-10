# 🧠 HashCache

`HashCache` is a simple and efficient **LRU (Least Recently Used) cache** library for Go with optional **Redis integration** for persistence and distributed caching.

---

## 🚀 Features

- ✅ In-memory LRU caching using Go’s `container/list`
- 🔁 Optional Redis fallback layer for persistence or shared access
- 🔒 Thread-safe with `sync.Mutex`
- 🧼 Auto-evicts least recently used entries beyond capacity
- 🧪 Simple, idiomatic API

---

## 📦 Installation

```bash
go get github.com/sapora1/hashcache
```

---

## 🛠️ Usage

### ✅ In-Memory Only

```go
package main

import (
	"fmt"
	"github.com/sapora1/hashcache"
)

func main() {
	cache := hashcache.NewCreateCache(2, nil)

	cache.Put("a", "apple")
	cache.Put("b", "banana")

	if val, ok := cache.Get("a"); ok {
		fmt.Println("Found:", val)
	}

	cache.Put("c", "cherry") // "b" will be evicted

	if _, ok := cache.Get("b"); !ok {
		fmt.Println("b has been evicted")
	}
}
```

---

### 🔁 With Redis (Optional)

```go
package main

import (
	"fmt"
	"github.com/sapora1/hashcache"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisOpts := &redis.Options{
		Addr: "localhost:6379",
	}

	cache := hashcache.NewCreateCache(2, redisOpts)

	cache.Put("x", "x-ray")

	val, ok := cache.Get("x")
	if ok {
		fmt.Println("Found in memory or Redis:", val)
	}
}
```

---

## 🔧 API Reference

### `NewCreateCache(capacity int, redisOptions *redis.Options) *Cache`

Creates a new LRU cache with optional Redis support.

- `capacity`: maximum number of entries in memory
- `redisOptions`: pass Redis options or `nil` to disable Redis

---

### `Put(key string, value interface{})`

Adds a new entry or updates an existing one.
- Stores in Redis if enabled.

---

### `Get(key string) (interface{}, bool)`

Fetches a value:
- First tries in-memory.
- Then falls back to Redis (if enabled).

---

### `Delete(key string)`

Deletes the key from in-memory cache.

---

### `Len() int`

Returns the number of in-memory entries.

---

## ⚡ Performance

| Operation | In-Memory LRU     | Redis (localhost) |
|-----------|-------------------|-------------------|
| `Get`     | ~100 ns – 5 µs    | ~100 – 500 µs     |
| `Put`     | ~200 ns – 10 µs   | ~200 – 800 µs     |

> Redis adds network and serialization overhead but enables multi-instance sharing.

---

## 🧪 Redis Setup (Local)

### Docker (recommended)

```bash
docker run --name redis-test -p 6379:6379 redis
```

### Ubuntu

```bash
sudo apt update
sudo apt install redis
sudo systemctl start redis
```

### macOS (Homebrew)

```bash
brew install redis
brew services start redis
```

---

## 🛡️ Thread Safety

This library uses `sync.Mutex` internally, so it is safe for concurrent access.

---

## 📁 Project Structure

```
hashcache/
├── cache.go        # LRU + Redis core implementation
├── go.mod
└── README.md       # 📖 You're here
```

---

## 🤝 Contributing

1. Fork this repo
2. Create a new branch: `git checkout -b feature-name`
3. Commit your changes
4. Push and open a PR

---

## 📄 License

MIT — feel free to use, fork, and improve.

---

## ✨ Author

Made with ❤️ by [Rahul](https://github.com/sapora1)

---
