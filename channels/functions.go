package main

import "fmt"
// import "builtin"


func BuffferedFibonacciSeries(limit int, buffChannel chan int){
  first, second := 0, 1
  buffChannel<-first
  buffChannel<-second
  sum := 0
  // fmt.Printf("%d %d", first, second)
  for i := 0; i < limit; i++{
      sum = first + second
      buffChannel<-sum
      // fmt.Printf(" %d", sum)
      first = second
      second = sum
  }
  close(buffChannel)
}

func UnbuffferedFibonacciSeries(limit int, buffChannel chan int){
  first, second := 0, 1
  buffChannel<-first
  buffChannel<-second
  sum := 0
  // fmt.Printf("%d %d", first, second)
  for i := 0; i < limit; i++{
      sum = first + second
      buffChannel<-sum
      // fmt.Printf(" %d", sum)
      first = second
      second = sum
  }
  close(buffChannel)
}

func main()  {
  limit := 10
  bufferedChannel := make(chan int, 5)
  unbufferedChannel := make(chan int)
  go BuffferedFibonacciSeries(limit, bufferedChannel)
  go UnbuffferedFibonacciSeries(limit, unbufferedChannel)
  for elem := range bufferedChannel {
        fmt.Printf("%d ", elem)
  }
  fmt.Printf("\n--------------------\n")
  for elem := range unbufferedChannel {
        fmt.Printf("%d ", elem)
  }
}
