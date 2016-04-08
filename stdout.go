package gosh

import "fmt"

func Stdout() {
	for _, pipe := range curPipes {
		wg.Add(1)
		go stdout(pipe)
	}
}

func stdout(pipe Pipe) {
	defer wg.Done()
	for data, moredata := <-pipe; moredata; {
		fmt.Print(data)
		data, moredata = <-pipe
	}
}
