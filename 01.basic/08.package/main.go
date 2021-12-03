package main

import (
	"log"

	"package/helpers"
)

func main() {
	log.Println("Hello, world!")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)

}
