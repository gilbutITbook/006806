package main

import (
	"fmt"
	"runtime"
)

type counter struct {
	i int64
}

// counter 값을 1씩 증가시킴
func (c *counter) increment() {
	c.i += 1
}

// counter 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	// 모든 CPU를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	// c.increment()를 실행하는 고루틴 1000개 실행
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()      // 카운터 값을 1 증가시킴
			done <- struct{}{} // done 채널에 완료 신호 전송
		}()
	}

	// 모든 고루틴이 완료될 때까지 대기
	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display() // c의 값 출력
}
