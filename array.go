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

	var slice = arr5[0:10]  //复制 arr5 数组的部分数据，从下标 0 开始到10结束
	fmt.Println(slice)

	var slice1 = arr5[:5]    //复制 arr 5 数据，省略了开始位置，默认为 0 开始，到 5 结束
	fmt.Println(slice1)

	var slice2 = arr5[5:]    //复制 arr 5 数据，从 5 开始，省略了结束位置，默认到 arr5 数组的最后一个位置
	fmt.Println(slice2)

	var make1 = make([]int , 5, 8)  //创建一个切片对象，长度默认为 5, 每次扩容 8 个长度
	fmt.Println(make1)

	var make2 = append(make1, 1,2,3,4,5,6,7)
	fmt.Printf("len = %d cap = %d %v\n", len(make2), cap(make2), make2)

  // len 获取切片的长度
  // cap 获取切片目前最多能存储多少元素

}
