package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// mapの初期化
	myMap := make(map[string]string)

	// 普通はこの初期化はしない
	var myMap2 map[string]string
	// myMap2["key"] = "value" と書くとpanicする
	// 以下のようにする必要が出るため、makeで初期化するのが一般的
	// if myMap2 == nil {
	// 	myMap2 = make(map[string]string)
	// }

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Buddy"
	log.Println(myMap)
	log.Println(myMap["dog"])
	log.Println(myMap2)

	// 構造体をMapに格納することもできる
	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	myMap3["me"] = me
	log.Println(myMap3)

	// その他のMapの重要な特徴
	// mapはポインタ型として扱われるので、値を変更すると、元のmapも変更される
	// mapのキーの並び順は、意図的にランダム化されるので、並び順は保証されてない
	// mapのvalueの型が不明な場合は、interface{}として扱うこともできる
	// ex) myMap[string]interface{}
	//

}
