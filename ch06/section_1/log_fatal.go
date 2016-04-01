package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
}
