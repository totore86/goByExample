package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond * 10)

			}
			wg.Done()
		}()
		fmt.Println(atomic.LoadUint64(&ops))

	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
