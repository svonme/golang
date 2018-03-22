package main

import "fmt"

func main()  {
  Break1:
    for {
      for i := 0; i < 10; i++ {
        fmt.Println("循环中的 i : ", i)
        if (i > 5) {
          break Break1 //结束整个标签的逻辑运算
        }
      }
    }

  fmt.Println("结束循环 Break1")


  Break2:
    for i := 0; i < 10; i++ {
      for j := 0; j < 100; j++ {
        fmt.Println("循环中的 j : ", j)
        if (i > 5) {
          break Break2 //结束整个标签的逻辑运算
        }
        if (j > 5) {
          fmt.Printf("continue Break2 第%d次\n", i + 1)
          continue Break2 //重新继续改标签的逻辑运算
        }
      }
    }

  fmt.Println("结束循环 Break2")

  for i := 0; i < 10; i++ {
    fmt.Println("goto i : ", i)
    if i > 5 {
      goto GOTO1
    }
  }

  GOTO1:
    fmt.Println("进入 GOTO 标签逻辑运算")
}
