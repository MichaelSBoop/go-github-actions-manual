package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func MaxInt(a, b int) int {
    if a >= b {
        return a
    }

    return b
}
