package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type PropertiesB map[string]map[string]string

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

	// load json for map test
	raw, err := ioutil.ReadFile("mapFile.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// unmarshall json to struct
	var objectB PropertiesB
	json.Unmarshal(raw, &objectB)

	// start counting for map
	startTime = makeTimestamp()

	// find in map
	for index, element := range objectB {

		if index == findWhat {

			var highest int64
			var highestValue string
			var i64 int64

			for i, v := range element {

				i64Prep, _ := strconv.ParseInt(i, 10, 64)
				i64 = i64Prep

				if i64 > highest {
					highest = i64
					highestValue = v
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
