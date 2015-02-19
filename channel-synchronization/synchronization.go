package main

import (
  "fmt"
  "time"
)

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}

func main(){

  done := make(chan bool, 1)
  go worker(done)

  // Block until we receive a notification from the worker on the channel.
  <- done //If you removed the line from this program, the program would exit before the worker even started.
}
