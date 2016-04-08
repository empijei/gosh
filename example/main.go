package main

import "github.com/RobClap/gosh"

func main() {
	gosh.Start()
	gosh.Cat("/tmp/tmp.txt")
	gosh.Stdout()
	gosh.End()
}
