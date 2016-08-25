package main

import (
	"xuexigo/gopl/chapter2/popcount"
	"os"
	"log"
	"fmt"
	"crypto/sha256"
	"encoding/binary"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("input 2 thing for sha256")
	}
	sha2561 := sha256.Sum256([]byte(os.Args[1]))
	sha2562 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%s\t%x\n%v\n", os.Args[1], sha2561, numberOfBits(sha2561[:]))
	fmt.Printf("%s\t%x\n%v\n", os.Args[2], sha2562, numberOfBits(sha2562[:]))
}

func numberOfBits(b []byte) (nums int) {
	if (len(b) % 8) != 0 {
		return -1
	}
	for i := 0; i < len(b) / 8; i ++ {
		ui64 := binary.LittleEndian.Uint64(b[i:8*(i+1)])
		nums += popcount.PopCount(ui64)
	}
	return
}
