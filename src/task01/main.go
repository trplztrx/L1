package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Weight int
	Height int
}

func (h Human) Speak() {
	fmt.Printf("My name is %s, i'm %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	ActionType string
}

func main() {
	action := Action{
		Human: Human{
			Name:   "Kostya",
			Age:    21,
			Weight: 70,
			Height: 180,
		},
		ActionType: "Speaking",
	}

	action.Speak()
	fmt.Println("Action Type:", action.ActionType)
}
