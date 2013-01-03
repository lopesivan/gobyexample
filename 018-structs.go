package main

import "fmt"

type person struct {
  name string
  age int
  age2 int
}

func main() {
  fmt.Println(person{"Bob", 20, 11})

  fmt.Println(person{name: "Alice", age: 30})

  fmt.Println(person{name: "Fred"})

  fmt.Println(&person{name: "Ann", age: 40})

  s := person{age: 50, name: "Sedkdkdkdkkdkdkdkdkkdkdkdkkdkdkdkkdkdkan"}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)
  fmt.Println(&sp.name)
  fmt.Println(&sp.age)
  fmt.Println(&sp.age2)
}
