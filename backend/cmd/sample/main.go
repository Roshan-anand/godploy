package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	// workers := 3
	wg := new(sync.WaitGroup)
	var id atomic.Int32

	for range 10 {
		wg.Go(func() {
			newID := id.Add(1)
			println("Generated ID:", newID)
		})
	}

	if wg.Wait(); true {
		println("All goroutines completed.")
	}
}
