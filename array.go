package main

import (
	"fmt"
)


func main() {
  var arr1 [10]int   //定义了一个长度为10的整型数组
  fmt.Println(arr1)

  // 定义一个长度为10的整型数组，并赋值默认值
  var arr2 = [10]int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(arr2)


  // 定义一个未声明长度的整型数组，并赋值默认值，长度由系统自动计算
  arr3 := []int{1,2,3,4,5}
  fmt.Println(arr3)

  // 效果同上，数组长度有系统自动计算
  arr4 := [...]int{1,2,3,4,5}
  fmt.Println(arr4)

  // 效果同上，但是只对数组下标为 19 的元素赋值，其余元素为系统默认值
  arr5 := [...]int{19: 11}
  arr5[10] = 100 //对单个元素赋值
  fmt.Println(arr5)
}
