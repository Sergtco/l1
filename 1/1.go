package main

import "fmt"

// Simple Human struct
type Human struct {
	Name     string
	Surname  string
	Age      int
	Strength int
}

// Raises Strength by amount of hours
func (h *Human) Train(hours int) {
	h.Strength += hours
}

// Generates default human
func NewHuman() Human {
	return Human{
		Name:     "John",
		Surname:  "Doe",
		Age:      22,
		Strength: 12,
	}
}

// Action embeds Human struct with it's fields and methods
type Action struct {
	Human
}

// Generates default action
func NewAction() Action {
	return Action{
		Human: NewHuman(),
	}
}

func main() {
	// Making action with default strength
	action := NewAction()
	fmt.Println("Current strength", action.Strength) // We can call fields of embedded struct.
	action.Train(3)                                  // And we can call methods of embedded struct
	fmt.Println("Current strength", action.Strength)
}
