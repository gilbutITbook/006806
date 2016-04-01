package main

import (
	"os"
	"runtime"
	"sync"
	"text/scanner"
)

type partial struct {
	token string
	scanner.Position
}

func mapper(path string, out chan<- partial) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	// 정상적인 파일이 아닌 경우 바로 리턴
	if info, err := file.Stat(); err != nil || info.Mode().IsDir() {
		return
	}

	var s scanner.Scanner
	s.Filename = path
	s.Init(file)

	// 파일의 모든 토큰을 스캔하여 out 채널로 전송
	tok := s.Scan()
	for tok != scanner.EOF {
		out <- partial{s.TokenText(), s.Pos()}
		tok = s.Scan()
	}
}

func runMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BUF_SIZE)
	go func() {
		for path := range paths {
			mapper(path, out)
		}
		close(out)
	}()
	return out
}

func runConcurrentMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BUF_SIZE)

	// mapper 작업을 CPU 코어 수 만큼 동시에 처리하도록 함
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				mapper(path, out)
			}
		}()
	}

	go func() {
		// 모든 mapper에서 작업을 완료할 때 까지 대기한 후 out 채널을 닫음
		wg.Wait()
		close(out)
	}()

	return out
}
