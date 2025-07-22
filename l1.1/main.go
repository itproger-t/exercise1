package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

func (h Human) GetName() string {
	return h.Name
}

func (h Human) GetAge() uint {
	return h.Age
}

type Action struct {
	Human
}

func (a Action) Meet() {
	fmt.Printf("Hi, my name is %s and i'm %d, what's your name?\n", a.GetName(), a.GetAge())
}

func main() {
	action := Action{
		Human{
			Name: "John",
			Age:  25,
		},
	}
	action.Meet()
}
