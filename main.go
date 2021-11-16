package main

import (
	"fmt"
	router "gee/routers"
	"log"
	"sync"
	"time"
)

func onlyForV2() router.HandlerFunc {
	return func(c *router.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))

	}
}

type student struct {
	Name string
	Age  string
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	// m.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}

	time.Sleep(time.Second)
}
