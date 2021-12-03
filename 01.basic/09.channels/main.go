package main

import (
	"channels/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	// Channelに値を送信する
	intChan <- randomNumber

}

func main() {
	// Channelはmakeで作成する
	intChan := make(chan int)

	// deferでchannelをcloseするのは良い習慣
	defer close(intChan)

	// goキーワードでgoroutineを起動する
	go CalculateValue(intChan)

	// Channelから値を受信する
	num := <-intChan
	log.Println(num)
}
