package goroutinepool

import (
	"sync"
	"reflect"
)

func returni(i int) int {
	return i
}

func sumAll(sumChan chan int) {
	var sum int
	for i := range sumChan {
		sum += i
	}
}

var goRoutinePoolLimit int = 100
var channelSize int = 100

func GoPool(counts int) {
	wg := sync.WaitGroup{}
	dataChan := make(chan int, channelSize)
	sumChan := make(chan int)
	go sumAll(sumChan)
	for i := 1; i <= goRoutinePoolLimit; i++ {
		wg.Add(1)
		go func() {
			var data int
			for data = range dataChan {
				sumChan <- returni(data)
			}
			wg.Done()
		}()
	}
	for j := 1; j <= counts; j++ {
		dataChan <- j
	}
	close(dataChan)
	wg.Wait()
	close(sumChan)
}

// 很明显如果counts取到大值的时候, process肯定就会爆掉了, 资源使用太多了. pool到可以避免这种麻烦.
// 但如果不使用pool, 而要避免process爆掉的情况, 那就肯定需要对可以用来运行function的goroutine的
// 总数做一个限制
func GoNoPool(counts int) {
	wg := sync.WaitGroup{}
	sumChan := make(chan int)
	go sumAll(sumChan)
	goRoutineLimit := 10000
	grlimtChan := make(chan struct{}, goRoutineLimit)
	for i := 1; i <= counts; i++ {
		wg.Add(1)
		grlimtChan <- struct{}{}
		go func(i int) {
			sumChan <- returni(i)
			wg.Done()
			<-grlimtChan
		}(i)
	}
	wg.Wait()
	close(sumChan)
	close(grlimtChan)
}

type reflectWorker struct {
	functor interface{}
	args    []reflect.Value
}

func GoPoolWithReflectWorker(counts int) {
	wg := sync.WaitGroup{}
	workerChan := make(chan reflectWorker, channelSize)
	sumChan := make(chan int)
	go sumAll(sumChan)

	for i := 1; i <= goRoutinePoolLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var worker reflectWorker
			for worker = range workerChan {
				sumChan <- int(reflect.ValueOf(worker.functor).Call(worker.args)[0].Int())
			}
		}()
	}
	for j := 1; j <= counts; j++ {

		workerChan <- reflectWorker{
			functor: returni,
			args: []reflect.Value{reflect.ValueOf(j)},
		}
	}
	close(workerChan)
	wg.Wait()
	close(sumChan)
}

func GoPoolWithRefReflectWorker(counts int) {
	wg := sync.WaitGroup{}
	workerChan := make(chan *reflectWorker, channelSize)
	sumChan := make(chan int)
	go sumAll(sumChan)

	for i := 1; i <= goRoutinePoolLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var worker *reflectWorker
			for worker = range workerChan {
				sumChan <- int(reflect.ValueOf(worker.functor).Call(worker.args)[0].Int())
			}
		}()
	}
	for j := 1; j <= counts; j++ {

		workerChan <- &reflectWorker{
			functor: returni,
			args: []reflect.Value{reflect.ValueOf(j)},
		}
	}
	close(workerChan)
	wg.Wait()
	close(sumChan)
}

type closureWorker struct {
	functor func()
}

func GoPoolWithClosureWorker(counts int) {
	wg := sync.WaitGroup{}
	workerChan := make(chan closureWorker, channelSize)
	sumChan := make(chan int)
	go sumAll(sumChan)
	for i := 1; i <= goRoutinePoolLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var worker closureWorker
			for worker = range workerChan {
				worker.functor()
			}
		}()
	}
	for j := 1; j <= counts; j++ {
		jj := j
		workerChan <- closureWorker{
			functor: func() {
				sumChan <- returni(jj)
			},
		}
	}
	close(workerChan)
	wg.Wait()
	close(sumChan)
}

func GoPoolWithRefClosureWorker(counts int) {
	wg := sync.WaitGroup{}
	workerChan := make(chan *closureWorker, channelSize)
	sumChan := make(chan int)
	go sumAll(sumChan)
	for i := 1; i <= goRoutinePoolLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var worker *closureWorker
			for worker = range workerChan {
				worker.functor()
			}
		}()
	}
	for j := 1; j <= counts; j++ {
		jj := j
		workerChan <- &closureWorker{
			functor: func() {
				sumChan <- returni(jj)
			},
		}
	}
	close(workerChan)
	wg.Wait()
	close(sumChan)
}
