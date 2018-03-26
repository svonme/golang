package main

import (
	"fmt"
)

func main()  {
  // 创建一个 map 对象
  map1 := make(map[int]string)
  // 赋值
  map1[1] = "hello wolrd"
  fmt.Println(map1)
  // 删除某个键值
  delete(map1, 1)

  // 复杂 map 结构
  map2 := make(map[int]map[int]string)
  // ok 可以判断该键值对是否存在
  map2_1, ok := map2[1]
  // 如果不存在，则重新初始化
  if !ok {
    map2[1] = make(map[int]string)
  }
  map2_1 = map2[1]
  map2_1[1] = "hello world"
  fmt.Println(map2_1[1])
  fmt.Println(map2)

  map3 := make(map[string]string)
  map3["name"] = "张三"
  map3["sex"] = "男"
  map3["age"] = "25"
  // 循环
  for key, value := range map3 {
    fmt.Println(key, value)
  }
}
