package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	chDone := make(chan struct{})

	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
		}
		log.Println("done")
		chDone <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-chDone

}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := copyBuffer(w, r); err != nil {
		log.Println(err)
	}
}

func copyBuffer(dst io.Writer, src io.Reader) (written int64, err error) {


	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			err = er
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}