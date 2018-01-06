package main

import "fmt"

type contactInfo struct {
  email string
  twitter string
}

type person struct {
  firstName string
  lastName string
  homeTown string
  nickname string
  contact contactInfo
}

func main() {
  // jamal := person{"Jamal", "O'Garro", "NYC"}
  // fmt.Printf("%+v", jamal)

  // var friend person
  // friend.firstName = "Cool"
  // friend.lastName = "Person"
  // friend.homeTown = "NYC"
  // fmt.Printf("%+v", friend)

  mj := person{
    firstName: "Michael",
    lastName: "Jordan",
    homeTown: "NC",
    nickname: "default",
    contact: contactInfo{
      email: "mike@bnike.com",
      twitter: "jumpman23",
    },
  }
  // fmt.Printf("%+v", mj)

  mj.sayName()

  // Pass reference of our struct to editName method
  (&mj).editName("Mike")

  mj.sayName()

  mjPointer := &mj
  mjPointer.setNickname("His Airness")
  mj.sayNickname()
}

func (p person) sayName() {
  fmt.Printf("%+v \n", p.firstName)
}

func (p person) sayNickname() {
  fmt.Printf("%+v \n", p.nickname)
}

func (p *person) editName(newName string) {
  // dereference the pointer
  (*p).firstName = newName
}

func (p *person) setNickname(n string) {
  (*p).nickname = n
}
