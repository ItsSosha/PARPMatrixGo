package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	fmt.Println("Sent done")
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done
	fmt.Println("Got done")
}
