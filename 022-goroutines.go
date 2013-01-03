package main

import "fmt"
import "time"

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
    time.Sleep(10000)
  }
}

func main() {
  f("direct")

  go f("goroutine")

  go func(msg string) {
    fmt.Println(msg)
  }("going")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
