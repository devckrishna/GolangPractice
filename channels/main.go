package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	numberOfTimes := 3
	go getNinjas(channel, numberOfTimes)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func getNinjas(channel chan string, numberOfTimes int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfTimes; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("you scored: ", score)
	}
	close(channel)
}
