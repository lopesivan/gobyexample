package main

import (
  "fmt"
  "math/rand"
  "runtime"
  "sync"
  "sync/atomic"
  "time"
)

func fakewait() {
  for x := 0; x < 1000000; x++ {
    w := rand.Intn(4)
    if w == 5 {
      fmt.Println("??")
    }
  }
}

func main() {
  var state = make(map[int]int)

  var mutex = &sync.Mutex{}

  var ops int64 = 0
  var wops int64 = 0

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

  for w := 0; w < 10; w++ {
    go func() {
      for {
        key := rand.Intn(5)
        mutex.Lock()
        state[key]++
        mutex.Unlock()
        atomic.AddInt64(&wops, 1)
        runtime.Gosched()
      }
    }()
  }

  time.Sleep(time.Second)

  opsFinal := atomic.LoadInt64(&ops)
  wopsFinal := atomic.LoadInt64(&wops)
  time.Sleep(time.Second)
  fakewait()
  sum := state[0]+state[1]+state[2]+state[3]+state[4]
  diff := sum - int(wopsFinal)
  fmt.Println("ops:", opsFinal)
  fmt.Println("wops:", wopsFinal)
  fmt.Println("sum:", sum)
  fmt.Println("diff:", diff)
  fmt.Println("state:", state)
  mutex.Lock()
  mutex.Unlock()
}
