package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(buffChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		buffChan <- i
		fmt.Printf("%d ", <-buffChan)
		time.Sleep(100 * time.Millisecond)
	}
	close(buffChan)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	buffChannel := make(chan int, 100)

	wg.Add(1)
	fmt.Println("Start of the concurrent run.")
	go generateNumbers(buffChannel, &wg)
	wg.Wait()
	fmt.Println("End of the concurrent run.")
}
