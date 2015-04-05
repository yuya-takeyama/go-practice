package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var times = 1000
var cnt uint64

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&cnt, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Count: %d\n", cnt)
}
