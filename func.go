package main

import (
	"fmt"
)

func main()  {
  func1()
  func2(10, 20)

  sum := Sum(5, 5)
  fmt.Printf("5 + 5 = %d \n", sum)

  num1 := Sum2(5, 8)
  fmt.Println(num1)


  num2, arr1 := func3(1,2,3,4,5)
  fmt.Println(num2, arr1)
}

// 无返回值
func func1()  {
  fmt.Println("无返回值函数")
}

// 带参函数
func func2(num1 int, num2 int)  {
  fmt.Printf("number1 : %d, number2 : %d\n", num1, num2)
}


// 有返回值，命名型返回值
func Sum(num1 int, num2 int) (sum int)  { //在定义返回值时已经隐性的根据参数创建了一个变量在函数内部
  sum = num1 + num2
  return  //这里可以省略返回的参数，在命名返回值时已经声明过需要返回的类型与名称，系统会默认使用声明函数中的参数
}


// 有返回值，ni匿名型返回值
func Sum2(num1 int, num2 int) int  { //如果只有一个返回值，可以身略括号，只写一个类型
  sum := num1 + num2
  return sum
}

// 任意多个参数, 多返回值
func func3(num int, array ...int) (int, []int)  { //如果只有一个返回值，可以身略括号，只写一个类型
  return num, array
}
