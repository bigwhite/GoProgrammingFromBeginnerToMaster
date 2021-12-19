package main

import (
	"sync"
	"time"
)

func goSleep() {
	time.Sleep(time.Second * 10)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			goSleep()
			wg.Done()
		}()
	}

	wg.Wait()
}
