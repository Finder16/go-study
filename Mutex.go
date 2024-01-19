package main

import (
  "fmt"
  "sync"
)

var wg = &sync.WaitGroup{}


func main(){
  myChannel := make(chan int)
  wg.Add(2)
  go runLoopSend(10, myChannel)
  go runLoopReceive(myChannel)
  wg.Wait()
}

func runLoopSend(n int, ch chan int) {
  defer wg.Done()
  for i := 0; i< n; i++{
    ch <- i
  }
close(ch)  
}

func runLoopReceive(ch chan int){
  defer wg.Done()
  for{
    i , ok := <- ch
    if !ok {
      break
    }
    fmt.Println("received value: " i)
  }
  
}
