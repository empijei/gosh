package gosh

import "net"

func Host2IP(in Pipe) (out Pipe) {
	out = NewPipe()
	addWait()
	go host2ip(in, out)
	return
}

func host2ip(in Pipe, out Pipe) {
	defer removeWait()
	defer close(out)
	for data, moredata := <-in; moredata; data, moredata = <-in {
		addresses, _ := net.LookupHost(data) //TODO handle err?
		for _, address := range addresses {
			out <- address
		}
	}
}

func IP2Host(in Pipe) (out Pipe) {
	out = NewPipe()
	addWait()
	go ip2host(in, out)
	return
}

func ip2host(in Pipe, out Pipe) {
	defer removeWait()
	defer close(out)
	for data, moredata := <-in; moredata; data, moredata = <-in {
		hosts, _ := net.LookupAddr(data) //TODO handle err?
		for _, host := range hosts {
			out <- host
		}
	}
}
