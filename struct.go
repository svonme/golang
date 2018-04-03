package main

import "fmt"

type person struct {
  Name string
  Age int
}

func main() {
  test3()
}

func test1 () {
  _person := person{}
  _person.Name = "张三"
  _person.Age = 20
  fmt.Println(_person)
}

func test2 () {
  _person := person{
    Name: "李四",
    Age: 25,
  }
  fmt.Println(_person)
}

func test3 () {
  _person := struct{
    Name string
    Age int
  }{
    Name: "王五",
    Age: 25,
  }
  fmt.Println(_person)
}
