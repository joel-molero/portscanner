package main

import (
	"fmt"
	"net"
	"sync"
)

func scan(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	//timeout := 2 * time.Second
	host := "scanme.nmap.org"

	addr := fmt.Sprintf("%s:%d", host, port)
	// conn, err := net.DialTimeout("tcp", addr, timeout)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("port %d closed (or filtered): %v\n", port, err)
		return
	}
	conn.Close()
	fmt.Printf("port %d open\n", port)
}

func main() {
	var wg sync.WaitGroup
	for port := 20; port <= 1024; port++ {
		wg.Add(1)
		go scan(port, &wg)
	}
	wg.Wait()
}
