package main

import (
	"fmt"
	"github.com/davidpoblador/gotrainz/train"
	"io"
	"log"
	"net/http"
)

var trains map[uint8]train.Train

func main() {
	trains = make(map[uint8]train.Train)

	http.HandleFunc("/train/list", ListTrains)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	NewTrain(4, "loco 2")
	fmt.Println(trains)

	NewTrain(1, "loco 1")
	fmt.Println(trains)

	t1 := trains[4]
	fmt.Println("t1 ", t1)
	t1.DirectionChange()
	fmt.Println("t1 ", t1)

	fmt.Println(trains[99])

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

// hello world, the web server
func ListTrains(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dummy response\n")
}

// Add a new train
func NewTrain(id uint8, description string) bool {

	if _, ok := trains[id]; ok == true {
		return false
	}

	train := new(train.Train)
	train.Description = description
	train.Forward = true
	train.Speed = 0

	// Store train
	trains[id] = *train

	return true
}
