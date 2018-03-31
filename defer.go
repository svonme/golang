package main

import "fmt"

/*
  它可以在函数返回前执行一些操作，最常用的就是打开一个资源（例如一个文件、数据库连接等）时就用defer延迟关闭改资源，以免引起内存泄漏
  执行顺序是逆序的，也就是先进后出的顺序
*/

func main()  {
  fmt.Println("start")
  fmt.Println("....")
  defer fmt.Println("stop")
}
