package main

import (
	"net"
	"log"
	"io"
	"time"
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
			log.Println(err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05.0000\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}