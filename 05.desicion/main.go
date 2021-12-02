package main

import "log"

func main() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("It's a cat")
	case "dog":
		log.Println("It's a dog")
	default:
		log.Println("Cat is something else")
	}
}
