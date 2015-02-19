package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

  var state = make(map[int]int)
  var mutex = &sync.Mutex{}

  var ops int64 = 0

  for r := 0; r < 100; r++ {
    go func() {
      total := 0
      for {
        key := rand.Intn(5)
        mutex.Lock()
        total += state[key]
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)

        runtime.Gosched()
      }
    }()
  }


    // We'll also start 10 goroutines to simulate writes,
    // using the same pattern we did for reads.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // Let the 10 goroutines work on the `state` and
    // `mutex` for a second.
    time.Sleep(time.Second)

    // Take and report a final operations count.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // With a final lock of `state`, show how it ended up.
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
