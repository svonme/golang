package main

import "fmt"

func main()  {
  test()
}

func test()  {
  fmt.Println("程序执行")
  defer func()  {
    err := recover()  //收取异常
    fmt.Println("捕获到的异常信息 : ", err)

    if err != nil {
      fmt.Println("异常处理完毕")
    }
  }()

  info := make(map[string]string)
  info["message"] = "程序因为xxx原因出现异常"
  panic(info)  //程序抛出异常
}
