package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func multiples(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	fmt.Println("Starting the concurrent run.")
	for n := 2; n <= 12; n++ {
		wg.Add(1)
		go multiples(n)
	}
	wg.Wait()
	fmt.Println("Concurrent run ended.")
}