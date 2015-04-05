package main

import (
	"fmt"
	"runtime"
	"sync"
)

var times = 1000
var cnt uint64

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			cnt++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Count: %d\n", cnt)
}
