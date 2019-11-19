package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup

	for i := 1; i <= 10; i++ {
		//fmt.Printf("-----%d", i)
		waitgroup.Add(1)
		go chanPrint(i, &waitgroup)
		waitgroup.Wait()
	}
}

func chanPrint(i int, waitgroup *sync.WaitGroup) {
	fmt.Println(i)
	waitgroup.Done()
}
