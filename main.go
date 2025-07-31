package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	timeout := 2 * time.Second
	host := "scanme.nmap.org"
	for port := 20; port <= 1024; port++ {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", addr, timeout)

		if err != nil {
			fmt.Printf("port %d closed (or filtered): %v\n", port, err)
			continue
		}
		conn.Close()
		fmt.Printf("port %d open\n", port)
		//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		//status, err := bufio.NewReader(conn).ReadString('\n')
	}
	//net.Dial("tcp", "scanme.nmap.org:")
}
