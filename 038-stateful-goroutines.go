package main

import (
  "fmt"
  "math/rand"
  "sync"
  "sync/atomic"
  "time"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  val int
  resp chan bool
}

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

  reads := make(chan *readOp)
  writes := make(chan *writeOp)
  state[0] = 0
  state[1] = 0
  state[2] = 0
  state[3] = 0
  state[4] = 0
  fmt.Println("state:", state)

  go func() {
    for {
      select {
      case read := <-reads:
        read.resp <- state[read.key]
      case write := <-writes:
        state[write.key]++
        atomic.AddInt64(&wops, 1)
        write.resp <- true
      }
    }
  }()

  for r := 0; r < 100; r++ {
    go func() {
      for {
        read := &readOp{
          key: rand.Intn(5),
          resp: make(chan int)}
        reads <- read
        <-read.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  for w := 0; w < 10; w++ {
    go func() {
      for {
        write := &writeOp{
          key: rand.Intn(5),
          val: rand.Intn(100),
          resp: make(chan bool)}
        writes <- write
        <-write.resp
      }
    }()
  }

  time.Sleep(time.Second * 10)

  opsFinal := atomic.LoadInt64(&ops)
  wopsFinal := atomic.LoadInt64(&wops)
  //fakewait()
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

