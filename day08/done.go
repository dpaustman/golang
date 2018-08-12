package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	//for {
	//	n, ok := <- c
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Printf("worker %d received %d\n", id, n)
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}

}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int,wg *sync.WaitGroup) worker {
	w := worker{
		in:  make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers  [10]worker
	for i :=0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}


	for i, worker := range workers {
		worker.in <- 'a' +i
		wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' +i
		wg.Add(1)
	}

	wg.Wait()
}


func main() {
	chanDemo()
}
