package main

import (
	"fmt"
)

func main()  {
  array := make([]int, 5)  //创建一个切片对象，初始长度为 5
  fmt.Println(array)

  array1 := make([]int, 5, 10)  //第三个参数为预留空间, 当容器内的空间使用完毕后，会自加 10 个空间
  fmt.Printf("长度为: %d, 空间为: %d", len(array1), cap(array1))
}
