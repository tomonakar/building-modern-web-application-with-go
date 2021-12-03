package main

import "log"

type myStruct struct {
	FirstName string
	LastName  string
	Phone     string
	Age       int
	Address   string
}

func (m *myStruct) getFullName() string {
	return m.FirstName + " " + m.LastName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"
	myVar.LastName = "Doe"

	myVar2 := myStruct{
		FirstName: "Anna",
		LastName:  "Smith",
	}

	log.Println(myVar.getFullName())
	log.Default().Println(myVar2.getFullName())

}
