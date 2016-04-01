package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		<-time.After(5 * time.Second)
		ticker.Stop()
	}()
	for now := range ticker.C {
		fmt.Printf("%v\n", now)
	}
}
