package main

//#include <unistd.h>
//void cgoSleep() { sleep(1000); }
import "C"
import (
	"sync"
)

func cgoSleep() {
	C.cgoSleep()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			cgoSleep()
			wg.Done()
		}()
	}

	wg.Wait()
}
