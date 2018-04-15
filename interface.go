package main

import "fmt"

type USB interface {
  Name() string
  Connect()
}

type PhoneConnecter struct {
  name string
}

func (pc PhoneConnecter) Name() string {
  return pc.name
}

func (pc PhoneConnecter) Connect() {
  fmt.Println("Connect: ", pc.name)
}

func main()  {
  var a USB
  a = PhoneConnecter{
    name: "手机",
  }

  a.Connect()
}
