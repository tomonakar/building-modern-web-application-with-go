package main

import "log"

func main() {
	// for loop
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// range loop: index, valueの順で取り出せる
	animals := []string{"cat", "dog", "bird"}
	for j, animal := range animals {
		log.Println(j, animal)
	}

	for _, animal2 := range animals {
		log.Println(animal2)
	}

	// map loop key, valueの順で取り出せる
	animals3 := make(map[string]string)
	animals3["dog"] = "bark"
	animals3["cat"] = "meow"
	animals3["bird"] = "tweet"

	for k, v := range animals3 {
		log.Println(k, v)
	}

	// 文字列のループ
	// 文字列はバイトのスライス(rune)に変換される
	var firstLine = "Once upoon, there was a brave princess who..."
	for i, v := range firstLine {
		log.Println(i, ":", v)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Doe", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@Jones.com", 34})
	users = append(users, User{"Sally", "Brown", "sally@brawn.com", 40})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 26})

	// 構造体のloop
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
