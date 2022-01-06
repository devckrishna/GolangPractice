package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mut sync.Mutex
var wrMut sync.RWMutex

func main() {
	// doIterations()
	readAndWrite()
}

func doIterations() {
	numberOfIteration := 1000
	for i := 0; i < numberOfIteration; i++ {
		go increaseCount()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(count)
}

func increaseCount() {
	mut.Lock()
	count++
	mut.Unlock()
}

func readAndWrite() {
	read()
	read()
	write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	wrMut.RLock()
	defer wrMut.RUnlock()
	fmt.Println("Read Locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading done")
}

func write() {
	wrMut.Lock()
	defer wrMut.Unlock()

	fmt.Println("write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Writing Done")
}
