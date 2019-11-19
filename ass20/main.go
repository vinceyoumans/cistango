package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type CustTime struct {
	RFC3339 time.Time
	CTT     CTT
}

type CTT struct {
	Date CTDate
	Time CTTIme
	Zone string
}

type CTTIme struct {
	CTHour int
	CTMin  int
}
type CTDate struct {
	CTYear  int
	CTMonth time.Month
	CTDay   int
}

func timeconvert(now01 time.Time, loc string) CTT {

	var ctt CTT

	location, err := time.LoadLocation(loc)
	if err != nil {
		fmt.Println(err)
	}
	now02 := now01.In(location)

	ctt.Date.CTYear = now02.Year()
	ctt.Date.CTMonth = now02.Month()
	ctt.Date.CTDay = now02.Day()

	ctt.Time.CTHour = now02.Hour()
	ctt.Time.CTMin = now02.Minute()

	ctt.Zone, _ = now02.Zone()

	return ctt

}

func main() {
	// struct
	//date, time and zone

	//  create a timestamp, and several TS in the future

	var ct CustTime
	var cts []CustTime

	now01 := time.Now()

	ct.RFC3339 = now01
	ct.CTT = timeconvert(now01, "EST")
	cts = append(cts, ct)

	ct.CTT = timeconvert(now01, "MST")
	cts = append(cts, ct)

	ct.CTT = timeconvert(now01, "UTC")
	cts = append(cts, ct)

	ct.CTT = timeconvert(now01, "CET")
	cts = append(cts, ct)

	data, err := json.MarshalIndent(cts, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println("========ccc=====================")

}
