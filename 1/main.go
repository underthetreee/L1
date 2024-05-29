package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h *Human) Say(phrase string) {
	fmt.Println(phrase)
}

type Action struct {
	Human
}

func main() {
	person := Action{
		Human{
			Name:   "Alex",
			Age:    25,
			Gender: "male",
		},
	}
	person.Say("hello")
}
