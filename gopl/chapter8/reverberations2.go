package main

import (
	"net"
	"log"
	"time"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	l, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go reverberation(conn, input.Text(), 2 * time.Second)
	}
}

func reverberation(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}