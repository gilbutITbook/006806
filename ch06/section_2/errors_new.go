package main

import (
	"errors"
	"fmt"
)

func main() {
	// 새 에러 생성
	errNotFound := errors.New("Not found error")
	fmt.Println("error: ", errNotFound)
	fmt.Println("error: ", errNotFound.Error())
}
