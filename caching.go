package caching

import(
    "time"
    "sync"
)

type cacheEntry struct{
    createdAt time.Time
    val []byte
}

type Cache struct{
     CacheMap map[string]cacheEntry
     Mutex sync.Mutex
 }

func CreateCache(expiry time.Duration) *Cache{
    var mu sync.Mutex
    cachemap := make(map[string]cacheEntry)
    cache := *Cache{CacheMap: cachemap, 
                  Mutex: mu,
              }
    ticker := time.NewTicker(5*time.second)
    go func(){
        for range ticker.C{
            cache.reapLoop(expiry)
        }
    }()

    defer ticker.Stop()
  return &cache 
}


func (c *Cache) Add(key string, value []byte) {
    c.Mutex.Lock()
    c.CacheMap[key] = cacheEntry{createdAt: time.Now(),
                                 val: value,
                             }
    c.Mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
    c.Mutex.Lock()
    value, exists := c.CacheMap[key]
    if !exists{
        return nil, false
    }
    c.Mutex.Unlock()
    return value, true
}

func (c *Cache) reapLoop(expiry time.Duration){
    c.Mutex.Lock()
    for key, cacheentry := range c.CacheMap{
        if time.Since(cacheentry.createdAt) > expiry{
            delete(c.CacheMap, key)
        }
    }
    c.Mutex.Unlock()
}
