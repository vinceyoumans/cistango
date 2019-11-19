package main

import (
	"fmt"
	"time"
)

//   this is kind of like the above examples,  but I wanted to demonstrate
// a select statement.  I did not remove the randomeness so the results will be different everytime.

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	for i := 1; i <= 10; i++ {
		go doChannel(i, c1, c2)
	}

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received from 1", msg1)

		case msg2 := <-c2:
			fmt.Println("received from 2", msg2)
		}

	}
	time.Sleep(1 * time.Second)

}

func doChannel(i int, c1 chan int, c2 chan int) {
	//  should be run everytime there is an i

	if i < 5 {
		c1 <- i
	} else {
		c2 <- i
	}
}
