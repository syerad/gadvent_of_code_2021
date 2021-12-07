package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Signals struct {
	Signals []int16 `json:"signals"`
}

func main() {
	sig := getAntennaSignals()

	fmt.Println("Increased measurements count " + strconv.Itoa(countIncreasedSignals(sig)))
	fmt.Println("Increased triplet count " + strconv.Itoa(countIncreasedSignalsTriplet(sig)))
}

func getAntennaSignals() []int16 {
	signals := readJsonInput()

	return signals.Signals
}

func countIncreasedSignals(signals []int16) int {
	var count int
	for i := 0; i < len(signals)-1; i++ {
		if signals[i] < signals[i+1] {
			count++
		}
	}
	return count
}

func countIncreasedSignalsTriplet(signals []int16) int {
	var count int

	for i := 0; i < len(signals)-3; i++ {
		countFirstTriplet := signals[i] + signals[i+1] + signals[i+2]
		countSecondTriplet := signals[i+1] + signals[i+2] + signals[i+3]

		if countFirstTriplet < countSecondTriplet {
			count++
		}
	}
	return count
}

func readJsonInput() Signals {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/1/day1.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened day1.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var signals Signals

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(byteValue, &signals); err != nil {
		log.Fatalf(err.Error())
	}

	return signals
}
