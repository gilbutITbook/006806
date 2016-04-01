package main

import (
	"bytes"
	"fmt"
	"sync"
	"text/scanner"
)

type summary struct {
	// 키: token
	// 값: map[string]int
	//           키: file path
	//           값: token count
	m map[string]map[string]int

	// 공유 데이터 m을 보호하기 위한 뮤텍스
	mu sync.Mutex
}

func reducer(token string, positions []scanner.Position) map[string]int {
	result := make(map[string]int)
	for _, p := range positions {
		result[p.Filename] += 1
	}
	return result
}

func (s summary) String() string {
	var buffer bytes.Buffer

	for token, value := range s.m {
		buffer.WriteString(fmt.Sprintf("Token: %s\n", token))
		total := 0
		for path, cnt := range value {
			if path == "" {
				continue
			}
			total += cnt
			buffer.WriteString(fmt.Sprintf("%8d %s ", cnt, path))
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("Total: %d\n\n", total))
	}
	return buffer.String()
}

func runReduce(tokenPositions intermediate) summary {
	s := summary{m: make(map[string]map[string]int)}
	for token, positions := range tokenPositions {
		s.mu.Lock() // m 값을 변경하는 부분(임계영역)을 뮤텍스로 잠금
		s.m[token] = reducer(token, positions)
		s.mu.Unlock() // m 값 변경 완료 후 뮤텍스 잠금 해제
	}
	return s
}

func runConcurrentReduce(in intermediate) summary {
	s := summary{m: make(map[string]map[string]int)}
	var wg sync.WaitGroup
	for token, value := range in {
		wg.Add(1)
		go func(token string, positions []scanner.Position) {
			defer wg.Done()
			s.mu.Lock() // m 값을 변경하는 부분(임계영역)을 뮤텍스로 잠금
			s.m[token] = reducer(token, positions)
			s.mu.Unlock() // m 값 변경 완료 후 뮤텍스 잠금 해제
		}(token, value)
	}
	wg.Wait()
	return s
}
