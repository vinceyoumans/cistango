package main

import (
	"fmt"
)

// GenNum ...
func GenNum() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		// ch <- n
		// return

		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
			//time.Sleep(100 * time.Millisecond)
		}

	}()
	return ch
}

func main() {
	nextnum := GenNum()
	for i := 1; i <= 10; i++ {
		fmt.Println(<-nextnum)
	}
	// fmt.Println(<-nextnum)
	// fmt.Println(<-nextnum)
	// fmt.Println(<-nextnum)
	close(nextnum)

}
