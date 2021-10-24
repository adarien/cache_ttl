# Реализация TTL (Time-To-Live) для кэша.

### Пример:

```go
package main

import (
	"fmt"
	"github.com/adarien/cache_ttl"
)

func main() {
	cache := cache.New()
	cache.Set("userID", 42, time.Second*1)
	cache.Set("userName", "adarien", time.Second*3)
	cache.Set("userStatus", true, time.Second*100)
	fmt.Println(cache.currentCache)

	userID, err := cache.Get("userID")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userID)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(cache.currentCache)

	userID, err = cache.Get("userID")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userID)
	}

	time.Sleep(time.Second * 2)

	fmt.Println(cache.currentCache)
}
```
