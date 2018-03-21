package main

import (
	"fmt"
)


func main() {

  for i, len := 1, 10; i < len; i++ {
    fmt.Println("i : ", i)
  }

  i := 1
  for true {
    fmt.Println("i : ", i)
    i++
    if i > 10 {
      break
    }
  }

}
