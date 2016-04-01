package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int64
	mu sync.Mutex // 공유데이터 i를 보호하기 위한 뮤텍스
}

// counter 값을 1씩 증가시킴
func (c *counter) increment() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계영역)을 뮤텍스로 잠금
	c.i += 1      // 공유데이터 변경
	c.mu.Unlock() // i 값 변경 완료 후 뮤텍스 잠금 해제
}

// counter의 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	// 모든 CPU를 사용하도록 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	// c.increment()를 실행하는 1000개의 고루틴 실행
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()      // 카운터 값을 1 증가시킴
			done <- struct{}{} // done 채널에 완료 신호 전송
		}()
	}

	// 모든 고루틴이 완료될 때 까지 대기
	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display() // c의 값 출력
}
