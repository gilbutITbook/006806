package main

import (
	"fmt"
	"strconv"
)

const (
	JOB_COUNT = 10
	BUF_SIZE  = 5
)

type Job struct{ name, log string }
type jobHandler func(Job) Job

func (j Job) String() string {
	return "job name: " + j.name + "\n[log]\n" + j.log
}

func pipe(handler jobHandler, in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			out <- handler(job)
		}
		close(out)
	}()
	return out
}

func main() {
	done := pipe(last, pipe(third, pipe(second, pipe(first, prepare()))))
	for d := range done {
		fmt.Println(d)
	}
}

// prepare
func prepare() <-chan Job {
	out := make(chan Job, BUF_SIZE)
	go func() {
		// prepare job
		for i := 0; i < JOB_COUNT; i++ {
			out <- Job{name: strconv.Itoa(i)}
		}
		close(out)
	}()
	return out
}

// job handlers
func first(job Job) Job {
	// first processing
	// ...
	job.log += "first stage\n"
	return job
}

func second(job Job) Job {
	// second processing
	// ...
	job.log += "second stage\n"
	return job
}

func third(job Job) Job {
	// third processing
	// ...
	job.log += "third stage\n"
	return job
}

func last(job Job) Job {
	// last processing
	// ...
	job.log += "last stage\n"
	return job
}
