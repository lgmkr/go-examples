package main

import (
  "fmt"
  "time"
)

func main() {

  c1 := make(chan string, 1)

  go func() {
    time.Sleep(time.Second * 2)
    c1 <- "result 1"
  }()

  select{
    case res := <-c1:
      fmt.Println(res)
    case <- time.After(time.Second * 1): // <-Time.After awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.
      fmt.Println("timeout 1")
  }

  c2 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "result 2"
  }()

  select{
    case res := <-c2:
      fmt.Println(res)
    case <- time.After(time.Second * 3):
      fmt.Println("timeout 2")
  }
}
