package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `[
		{
			"first_name": "John",
			"last_name": "Doe",
			"hair_color": "Brown",
			"has_dog": true
		},
		{
			"first_name": "Jane",
			"last_name": "Doe",
			"hair_color": "Blonde",
			"has_dog": false
		},
		{
			"first_name": "Jack",
			"last_name": "Doe",
			"hair_color": "Black",
			"has_dog": true
		}
	]`

	var unmarshalled []Person

	//	JSONデータを読み込むときは、Unmarshal関数を使う
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println(unmarshalled)

	// JSONデータを書き出すときは、Marshal関数を使う
	var mySlice []Person

	var m1 Person
	m1.FirstName = "John"
	m1.LastName = "Doe"
	m1.HairColor = "Brown"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName = "Alice"
	m1.LastName = "Poe"
	m1.HairColor = "Brown"
	m1.HasDog = false

	mySlice = append(mySlice, m2)

	// インデント付きのJSONを生成するには、MarshalIndent関数を使う
	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("Error uarshalling json", err)
	}

	log.Println(string(newJson))
}
