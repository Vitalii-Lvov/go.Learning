package main

import (
	"fmt"
	"sync"
)

func main() {

	testLoop()

	var wg sync.WaitGroup
	wg.Add(1)

	go testLoop()

	wg.Wait()
}

func testLoop() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

}
