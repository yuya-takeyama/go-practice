package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

// Goroutine の数を数えておく
var gCnt = 0

// Goroutine の ID 発番
var gId = 1

// ループを実行する回数の設定
var loopCount = 500

func main() {
	// プロファイラを起動してブラウザで確認できるようにしておく
	// http://localhost:6060/debug/pprof/goroutine?debug=2 とかで Goroutine の様子を確認できる
	runtime.SetBlockProfileRate(1)
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(loopCount)

	c := make(chan bool, 10)
	for i := 1; i <= loopCount; i++ {
		c <- true
		go func() {
			mutex.Lock()
			myGid := gId
			fmt.Printf("%d is began\n", gId)
			gId++
			gCnt++
			updateGoroutineCount()
			mutex.Unlock()

			time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
			defer func() {
				mutex.Lock()
				fmt.Printf("%d is finished\n", myGid)
				gCnt--
				updateGoroutineCount()
				mutex.Unlock()

				wg.Done()
				<-c
			}()
		}()
	}

	wg.Wait()
}

func updateGoroutineCount() {
	fmt.Printf("%d goroutines exist\n", gCnt)
}
