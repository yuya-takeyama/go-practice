package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const times = 10000

func useIncrementOperator() uint32 {
	var cnt uint32
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			cnt++
			wg.Done()
		}()
	}

	wg.Wait()

	return cnt
}

func useAtomicAddUint32() uint32 {
	var cnt uint32
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint32(&cnt, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	return cnt
}

func useSyncMutexLock() uint32 {
	var cnt uint32
	var wg sync.WaitGroup
	mu := new(sync.Mutex)

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt++
			wg.Done()
		}()
	}

	wg.Wait()

	return cnt
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	fmt.Printf("useIncrementOperator(): %d\n", useIncrementOperator())
	fmt.Printf("useAtomicAddUint32(): %d\n", useAtomicAddUint32())
	fmt.Printf("useSyncMutexLock(): %d\n", useSyncMutexLock())
}
