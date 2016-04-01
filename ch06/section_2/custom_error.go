package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type SqrtError struct {
	time    time.Time // 에러가 발생한 시간
	value   float64   // 에러를 발생시킨 값
	message string    // 에러 메시지
}

// error 인터페이스에 정의된 Error() 메서드 구현
func (e SqrtError) Error() string {
	return fmt.Sprintf("[%v]ERROR - %s(value: %g)", e.time, e.message, e.value)
}
func Sqrt(f float64) (float64, error) {
	// 매개변수로 전달된 값이 유효한 값이 아닐 때 SqrtError를 반환
	if f < 0 {
		return 0, SqrtError{time: time.Now(), value: f, message: "음수는 사용할 수 없습니다"}
	}
	if math.IsInf(f, 1) {
		return 0, SqrtError{time: time.Now(), value: f, message: "무한대 값은 사용할 수 없습니다"}
	}
	if math.IsNaN(f) {
		return 0, SqrtError{time: time.Now(), value: f, message: "잘못된 수 입니다"}
	}
	// 정상 처리 결과 반환
	return math.Sqrt(f), nil
}

func main() {
	v, err := Sqrt(9) // 9의 제곱근. 정상 동작
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	v, err = Sqrt(-9) // -1의 제곱근. 에러 발생
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
