package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}


var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}


func main() {
	start := time.Now()
	runtime.GOMAXPROCS(16)
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	elapsed := time.Since(start)
    fmt.Printf("Concurrent execution took %s\n", elapsed)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(16))
}

