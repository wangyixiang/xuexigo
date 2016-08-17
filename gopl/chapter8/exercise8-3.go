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
	// this one just works for reverberations1 but no reverberations2, why?

	// now I come to answer this, the reverberations2 use goroutine to serve
	// every new request parsing from input.Scan(), but not like reverberations1
	// which serve every new request after finishing serving the previous
	// request, so when the ^D request come, reverberation1 will finish all the
	// requests before it, but reverberation2 will return false from input.Scan()
	// and quit the handleConn then quit the main goroutine, that's why all
	// remaining goroutines will terminate either, and "exercise8-3" will quit
	// immediately even it just close the write half of TCP connection but it
	// still get error from io.Copy.

	// exercise8-4 give a chance to add synchronization in the reverberations2
	// to enable exercise8-3 work normally.
	tcpConn := conn.(*net.TCPConn)
	tcpConn.CloseWrite()
	<-chDone
	conn.Close()

}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := copyBuffer(w, r); err != nil {
		log.Println(err)
	}
}

func copyBuffer(dst io.Writer, src io.Reader) (written int64, err error) {

	buf := make([]byte, 32 * 1024)
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