package main

import (
	"fmt"
	// "time"
	"sync"
)

// func generateNumbers(buffChan chan int, wg *sync.WaitGroup){
//   for i := 0; i < 100 ; i++{
//     buffChan<-i
//     time.Sleep(100 * time.Millisecond)
//   }
//   close(buffChan)
//   wg.Done()
// }

func forLoop(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// buffChannel := make(chan int, 100)

	wg.Add(1)
	fmt.Println("Start of the concurrent run.")
	go forLoop(&wg)
	wg.Wait()
	fmt.Println("End of the concurrent run.")

	// fmt.Println("Start of the buffered channel.")
	// wg.Add(1)
	// go generateNumbers(buffChannel, &wg)
	// // fmt.Print(<-buffChannel)
	// for elem := range buffChannel {
	//       fmt.Printf("%d ", elem)
	// }
	// wg.Wait()
	// fmt.Println("End of the buffered channel.")
}
