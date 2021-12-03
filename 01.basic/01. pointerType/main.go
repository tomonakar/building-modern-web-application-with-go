package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is:", myString)
	changeUsingPointer(&myString)
	log.Println("myString is:", myString)
}

func changeUsingPointer(s *string) {
	log.Println("上の*は、ポインタ型であることを示す")
	log.Println("-------------------------------")
	log.Println("ポインタ変数のsは、アドレスがセットされている", s)
	log.Println("ポインタ変数の*s は、ポインタ変数の指すメモリアドレスの実データ", *s)
	log.Println("ポインタ変数の&s は、変数がメモリ上で確保された際のアドレス", &s)
	log.Println("-------------------------------")

	newValue := "Red"
	*s = newValue
}
