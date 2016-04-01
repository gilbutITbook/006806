package main

import "fmt"

const (
	// iota를 사용한 비트연산의 결과값으로 상수 선언
	Running = 1 << iota // 1 << 0 == 1
	Waiting             // 1 << 1 == 2
	Send                // 1 << 2 == 4
	Receive             // 1 << 3 == 8
)

func main() {
	// OR 연산자(|)로 stat 변수 생성
	stat := Running | Send
	display(stat)
}

func display(stat int) {
	// AND 연산자(&)로 stat에 포함된 비트값의 상태 출력
	if stat&Running == Running {
		fmt.Println("Running")
	}
	if stat&Waiting == Waiting {
		fmt.Println("Waiting")
	}
	if stat&Send == Send {
		fmt.Println("Send")
	}
	if stat&Receive == Receive {
		fmt.Println("Receive")
	}
}
