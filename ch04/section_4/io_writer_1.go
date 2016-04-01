package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func handle(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func main() {
	msg := []string{"This", "is", "an", "example", "of", "io.Writer"}
	for _, s := range msg {
		time.Sleep(100 * time.Millisecond)
		handle(os.Stdout, s)
	}
}
