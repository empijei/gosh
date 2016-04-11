package gosh

import "sync"

type Job func(string) string

func Parallel(in Pipe, job Job, j int) (out Pipe) {
	addWait()
	out = NewPipe()
	go parallel(in, job, j, out)
	return
}

func parallel(in Pipe, job Job, j int, out Pipe) {
	defer removeWait()
	defer close(out)
	var subwg sync.WaitGroup
	for i := 0; i < j; i++ {
		subwg.Add(1)
		go doJob(in, job, out, subwg)
	}
	subwg.Wait()
}

func doJob(in Pipe, job Job, out Pipe, subwg sync.WaitGroup) {
	defer subwg.Done()
	for data, moredata := <-in; moredata; data, moredata = <-in {
		out <- job(data)
	}
}
