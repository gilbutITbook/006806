package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	done := process(quit)
	timeout := time.After(1 * time.Second)

	select {
	case d := <-done:
		fmt.Println(d)
	case <-timeout:
		fmt.Println("Time out!")
		quit <- struct{}{}
	}
}

func process(quit <-chan struct{}) chan string {
	done := make(chan string)
	go func() {
		go func() {
			time.Sleep(10 * time.Second) // heavy job

			done <- "Complete!"
		}()

		<-quit
		return
	}()
	return done
}
