package gosh

import "fmt"

func Stdout(in Pipe) {
	addWait()
	go stdout(in)
}
func stdout(in Pipe) {
	defer removeWait()
	for data, moredata := <-in; moredata; data, moredata = <-in {
		fmt.Println(data)
	}
}
