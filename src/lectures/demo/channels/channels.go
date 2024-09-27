package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOK
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOK
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {

	// job
	jobs := make(chan Job, 50)
	// results
	results := make(chan Result, 50)
	// control
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Microsecond):
			fmt.Println("time out")
			control <- DoExit
			<-control
			fmt.Println("Program exit")
			return
		}
	}

}
