package main

import (
	fmt "fmt"
  "strconv"
  "reflect"
)

const (
  B float64 = 1 << (iota * 10)
  KB
  MB
  GB
  TB
  PB
)

func main() {
	var number int //变量的声明
	number = 20    //变量的赋值
	fmt.Println("先声明，后赋值", number)

	var count int = 50 //变量声明的同时并赋值
	fmt.Println("声明的同时并赋值", count)

	var name = "小明" //声明变量时不指定类型，由系统盘点类型
	fmt.Println("省略类型，由系统判断类型", name)

	age := 22 //声明变量的简写格式
	fmt.Println("简写", age)


  a := 65
  stra := strconv.Itoa(a)  //将数字转为字符串
  fmt.Println(stra, reflect.TypeOf(stra))

  b, _ := strconv.Atoi(stra)  //字符串转为数字
  fmt.Println(b, reflect.TypeOf(b))

  fmt.Println(B, KB, GB, PB)


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

  c := 5
  switch true {
  case c > 0:
    fmt.Println("c > 0")
    fallthrough //继续检查符合条件的语句
  case c > 1:
    fmt.Println("c > 1")
  case c > 2:
    fmt.Println("c > 2")
  default:
    fmt.Println("default")
  }

}
