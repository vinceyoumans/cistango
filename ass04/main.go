package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
)

type basket []struct {
	Basket struct {
		Account     string `json:"account"`
		ST01Metrics struct {
			SubmittionTs int    `json:"submittion_ts"`
			AccountID    string `json:"Account_id"`
		} `json:"ST01_Metrics"`
		Items []struct {
			Tag string `json:"tag"`
			Obj string `json:"obj"`
		} `json:"items"`
	} `json:"basket"`
}

var wg = sync.WaitGroup{}

func main() {

	var B basket

	jsonFile, err := os.Open("basketA01.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &B)

	ilen := len(B[0].Basket.Items)

	// I have an 8 core i7, so I should have 8 threads running
	fmt.Println("===========================================")
	fmt.Printf("AAA  Threads : %v \n", runtime.GOMAXPROCS(-1))
	fmt.Println("===========================================")

	fmt.Println("===========================================")
	runtime.GOMAXPROCS(100)
	fmt.Printf("BBB  Threads : %v \n", runtime.GOMAXPROCS(-1))
	fmt.Println("===========================================")
	// Changing MAXPROCS is a means of tunning the app

	// Part 1...  Print out original Basket
	// the Tags are not Modified
	fmt.Println("=====AAA  Unmodified Basket  ======")
	for i := 0; i < ilen; i++ {
		fmt.Printf("===  %d - %v \n", i, B[0].Basket.Items[i].Tag)
	}
	fmt.Println("=====AAA  Unmodified Basket  ======")

	// Part 2...  GoRoutines in WG to modify Basket Items
	// Assuming that a single basket has X number of items.
	// using GoRoutine, pass a pointer the the Basket struct
	// and modify it using a pointer.
	// Modifiers are done in paral using Go Routines
	// When the entire Wait group is Modified, the wg.Done() is called which release main thread to continue.
	fmt.Println("=====BBB  Begining Modification  ======")
	for i := 0; i < ilen; i++ {
		//fmt.Printf("===  %d - %v \n", i, B[0].Basket.Items[i].Tag)
		wg.Add(1)
		go Modified(i, &B[0].Basket.Items[i].Tag)
	}
	wg.Wait()
	fmt.Println("=====BBB  Ending Modification  ======")

	// Third Part of Example...
	// After all of the comutation is finsihed on this BASKET
	// The Tags should be Modified.  Or the Items should all be Modified.
	fmt.Println("=====CCC  Modified Basket  ======")
	for i := 0; i < ilen; i++ {
		fmt.Printf("===  %d - %v \n", i, B[0].Basket.Items[i].Tag)
	}
	fmt.Println("=====CCC  Modified Basket  ======")

}

func Modified(tt int, bb *string) {
	fmt.Println(tt)
	// fmt.Println(bb[0].Basket.Items[tt].Tag)
	fmt.Println(*bb)
	*bb = *bb + "-MODIFIED"
	wg.Done()
}

