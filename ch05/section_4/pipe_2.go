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

func main() {
	done := doLast(doThird(doSecond(doFirst(prepare()))))
	for d := range done {
		fmt.Println(d)
	}
}

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

func doFirst(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			// first processing
			// ...
			job.log += "first stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doSecond(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			// second processing
			// ...
			job.log += "second stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doThird(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			// third processing
			// ...
			job.log += "third stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doLast(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			// last processing
			// ...
			job.log += "last stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}
