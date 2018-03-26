package main

import (
	"fmt"
)

func main()  {
  arr1 := []int{1,2,3,4}
  arr2 := []int{5,6,7,8,11,22,33,44}

  //把 arr2 数组的内容复制到 arr1 中
  //复制时，只会复制 arr1 长度的内容，arr1 数组只有 4 个元素，arr2 数组有 8 个元素，只会把 arr2 数组中前 4 个元素复制到 arr1 中
  copy(arr1, arr2)
  fmt.Println(arr1, arr2)
  copy(arr1[2:], arr2[6:]) //只替换部分内容
  fmt.Println(arr1, arr2)
}
