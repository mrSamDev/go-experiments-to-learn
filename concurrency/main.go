package main

import (
	"time"
)

func slowGreet(doneVar chan bool, text string) {
	time.Sleep(3 * time.Second)
	println("Slow........... Hello......, World!", text)
	doneVar <- true
	close(doneVar)
}

func main() {
	done := make(chan bool)
	go slowGreet(done, "1")
	go slowGreet(done, "2")
	go slowGreet(done, "3")
	go slowGreet(done, "3")
	go slowGreet(done, "4")
	go slowGreet(done, "5")
	go slowGreet(done, "6")

	for range done {

	}

}
