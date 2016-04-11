package gosh

import "sync"

type Pipe chan string
type PipeBlock func(interface{}) Pipe

var bufSize = 1024
var wg sync.WaitGroup

func NewPipe(sizes ...int) Pipe {
	if len(sizes) == 0 {
		return make(Pipe, bufSize)
	} else {
		return make(Pipe, sizes[0])
	}
}
func End() {
	wg.Wait()
}

func addWait() {
	wg.Add(1)
}

func removeWait() {
	wg.Done()
}

func (pipe Pipe) ToStdout() {
	//TODO
}
func (pipe Pipe) ToFile(filename string) {
	//TODO
}
