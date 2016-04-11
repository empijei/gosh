package main

import "github.com/RobClap/gosh"

func main() {
	defer gosh.End()
	pipe := gosh.Cat("/tmp/addresses.txt")
	pipe = gosh.Host2IP(pipe)
	gosh.Stdout(pipe)
}
