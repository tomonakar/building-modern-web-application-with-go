package main

import "fmt"

type Animal interface {
	Say() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Fido",
		Breed: "Mixed",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "George",
		Color:         "Brown",
		NumberOfTeeth: 30,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Say(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Say() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Say() string {
	return "ughhhh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
