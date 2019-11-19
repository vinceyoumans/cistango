package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testbasket []struct {
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

func TestModified(t *testing.T) {
	var B testbasket

	jsonFile, err := os.Open("basketA01.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &B)

	ilen := len(B[0].Basket.Items)

	fmt.Println("=====AAA  Unmodified Basket  ======")
	for i := 0; i < ilen; i++ {
		wg.Add(1)
		a := B[0].Basket.Items[i].Tag
		a = a + "-MODIFIED"
		Modified(i, &B[0].Basket.Items[i].Tag)
		b := B[0].Basket.Items[i].Tag
		assert.Equal(t, a, b, "The two words should be the same.")
		fmt.Println(b)
	}
	wg.Wait()

	//===========================================
	// The following should fail

	fmt.Println("=====BBB  These tests should fail  ======")
	for i := 0; i < ilen; i++ {
		wg.Add(1)
		a := B[0].Basket.Items[i].Tag
		a = a + "-MODIFied again"
		Modified(i, &B[0].Basket.Items[i].Tag)
		b := B[0].Basket.Items[i].Tag
		assert.Equal(t, a, b, "The two words should be the same.")
		fmt.Println(b)
	}
	wg.Wait()
	fmt.Println("=====BBB  Unmodified Basket  ======")

}