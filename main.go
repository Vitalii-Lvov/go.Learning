package main

import (
	"fmt"
	"sync"
)

func main() {

	testLoop(nil)

	var wg sync.WaitGroup
	wg.Add(1)

	go testLoop(&wg)

	prn("waiting began")

	wg.Wait()

	prn("waitnig end")

}

func testLoop(wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	if wg != nil {
		wg.Done()
	}

}

func prn(msg any) {
	fmt.Println(msg)
}
