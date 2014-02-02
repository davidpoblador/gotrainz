package main

import (
	//	"encoding/json"
	"fmt"
	. "github.com/davidpoblador/gotrainz/train"
	"io"
	"log"
	"net/http"
)

var trains map[uint8]*Train

func main() {
	trains = make(map[uint8]*Train)

	http.HandleFunc("/train/list", ListTrains)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	//AddTrain(4, "loco 2")
	//fmt.Println(trains)
	//fmt.Println(*trains[4])
	//trains[4].DirectionChange()
	//AddTrain(1, "loco 1")
	////	fmt.Println(trains)
	//fmt.Println(*trains[4])
	//	fmt.Println(trains)
	// Test direction change
	//trains[1].DirectionChange()
	//fmt.Println(trains)

	// Test speed change
	//trains[1].SpeedChange(40)
	//fmt.Println(trains)

	// Test relative speed change
	//trains[1].RelativeSpeedChange(-20)
	//fmt.Println(trains)
}

// Delete a train from the system
func DelTrain(id uint8) {
	delete(trains, id)
}

// Add a new train
func AddTrain(id uint8, description string) {
	if _, ok := trains[id]; ok == true {
		return
	}
	// Store train
	a := NewTrain(description)

	trains[id] = a

	//	trains[id].fmt.Println(a)
}

// HTTP API

// List Trains
func ListTrains(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test\n")
	fmt.Println("")
}
