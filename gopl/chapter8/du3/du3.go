package main

import (
	"flag"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	"time"
	"sync"
	"runtime/pprof"
)

type countingSemaphore struct {

}

func walkDir(dir string, fileSize chan <- int64, wg *sync.WaitGroup, cS chan countingSemaphore) {
	dirs := dEntries(dir, cS)
	for _, entry := range dirs {
		if entry.IsDir() && (entry.Mode() & os.ModeSymlink == 0) {
			wg.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), fileSize, wg, cS)
			continue
		}
		fileSize <- entry.Size()
	}
	wg.Done()
}

func dEntries(dir string, cS chan countingSemaphore) []os.FileInfo {
	cS <- struct{}{}
	defer func() {
		<-cS
	}()
	results, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("du3: %v\n", err)
		return nil
	}
	return results
}

func main() {
	var verbose = flag.Bool("v", false, "show spin when waiting.")
	var cpuProf = flag.String("cpuprof", "", "write cpu profile to `file`")
	flag.Parse()
	if *cpuProf != "" {
		f, err := os.Create(*cpuProf)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}
	roots := flag.Args()
	var tick  <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Millisecond * 500)
	}

	if len(roots) == 0 {
		roots = append(roots, ".")
	}
	chFileSize := make(chan int64)
	cS := make(chan countingSemaphore, 20)
	fmt.Println(time.Now())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, root := range roots {
			wg.Add(1)
			go walkDir(root, chFileSize, wg, cS)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(chFileSize)
	}()

	nCount := 0
	nSize := int64(0)

	str := `-\|/`
	i := 0
	loop:
	for {
		select {
		case fSize, ok := <-chFileSize:
			if !ok {
				break loop
			}
			nCount += 1
			nSize += fSize
		case <-tick:
			fmt.Printf("\r(%c)%d files %dbytes %fG", str[i], nCount, nSize, float64(nSize) / 1e9)
			i = (i + 1) % 4
		}

	}
	fmt.Printf("\r%d files %dbytes %fG\n", nCount, nSize, float64(nSize) / 1e9)
	fmt.Println(time.Now())
}
