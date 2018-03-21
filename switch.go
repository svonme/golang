package main

import (
	"fmt"
)

func main() {
  number := 5
  switch true {
  case number > 0:
    fmt.Println("number > 0")
    fallthrough //继续检查符合条件的语句
  case number > 1:
    fmt.Println("number > 1")
  case number > 2:
    fmt.Println("number > 2")
  default:
    fmt.Println("default")
  }

}
