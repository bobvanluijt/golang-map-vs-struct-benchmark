package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type PropertiesA []struct {
	History []struct {
		Timestamp int64  `json:"timestamp"`
		Value     string `json:"value"`
	} `json:"history"`
	Property string `json:"property"`
}

type propertiesB map[string]map[string]string

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}

func main() {

	// EDIT FOR TESTS
	// WHAT SHOULD BE FOUND?
	findWhat := os.Args[1]

	// define time stamps
	var startTime int64
	var endTime int64

	// load json for struct test
	raw, err := ioutil.ReadFile("./structFile.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// unmarshall json to struct
	var objectA PropertiesA
	json.Unmarshal(raw, &objectA)

	// start counting for struct
	startTime = makeTimestamp()

	// find in struct
	for _, element := range objectA {
		if element.Property == findWhat {
			var highest int64
			var highestValue string
			for _, v := range element.History {
				if v.Timestamp > highest {
					highest = v.Timestamp
					highestValue = v.Value
				}
			}
			fmt.Println("STRUCT QUERY FOUND: ", highestValue)
		}
	}

	// end counting
	endTime = makeTimestamp()

	// show first result for struct
	fmt.Println("STRUCT QUERY TOOK: ", endTime-startTime)

	// load json for map test
	raw, err = ioutil.ReadFile("./mapFile.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// start counting for map
	startTime = makeTimestamp()

	// find in map
	for _, element := range objectA {

		if element.Property == findWhat {
			var highest int64
			var highestValue string
			for _, v := range element.History {
				if v.Timestamp > highest {
					highest = v.Timestamp
					highestValue = v.Value
				}
			}
			fmt.Println("MAP QUERY FOUND: ", highestValue)
		}

	}

	// end counting map
	endTime = makeTimestamp()

	// show first result for map
	fmt.Println("MAP QUERY TOOK: ", endTime-startTime)

}
