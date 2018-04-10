package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func main()  {
  person := Person{
    Name: "小明",
    Age: 20,
  }

  person.say()
}


func (person *Person) say() {
  fmt.Printf("我是%s， 我今年%d岁\n", person.Name, person.Age)
}
