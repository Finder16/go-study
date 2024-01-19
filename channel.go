package main

import (
  "fmt"
  "time"
)

func runLoopSend(n int , ch chan int){
  for i := 0; i<n; i++{
    ch <- i
    }
close()
  }
func runLoopReceive(ch chan int){
  for{
  i, ok :=<-ch
  if !ok{
    break
  }
    fmt.Println("received data",i)
    }

}

func main(){
  mycannel := make(chan int)
  go runLoopSend(10, mycannel)
  go runLoopSend(10, mycannel)
  time.Sleep(2 * time.Second)
}
