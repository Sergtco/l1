package main

import "fmt"

type NewInter interface { // New interface which should be implemented by Adaptee
	NewFunction()
}

type Adapter struct { // Adapter is an adapter to our old type
	a *Adaptee
}

func (a *Adapter) NewFunction() { // Adapter implements NewInter
	a.a.OldFunction()
}

// we cannot touch Adaptee
type Adaptee struct { // Old type which is using old interface
}

func (a *Adaptee) OldFunction() { // Adaptee cannot be of NewInter interface
	fmt.Println("Old function has been called!")
}

func main() {
	fmt.Println("Made new feature!")
	var feature NewInter = &Adapter{a: &Adaptee{}}
	feature.NewFunction()
	fmt.Println("Oops, it utilizes old code!")
}
