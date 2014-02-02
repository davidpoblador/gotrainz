package main

import (
	"fmt"
	"math"
)

type Train struct {
	id          uint8
	description string
	forward     bool
	speed       uint8
}

var trains map[uint8]Train

func (t *Train) DirectionChange() {
	t.forward = !t.forward
}

func (t *Train) Stop() {
	t.SpeedChange(0)
}

func (t *Train) SpeedChange(speed uint8) {
	// Actual stop code
	// ...
	// Change speed
	t.speed = speed
}

func (t *Train) RelativeSpeedChange(speed int) uint8 {
	newspeed := NormalizeSpeed(int(t.speed) + speed)

	t.SpeedChange(newspeed)
	return newspeed
}

func NormalizeSpeed(speed int) uint8 {
	return uint8(math.Min((math.Max(float64(speed), 0)), math.MaxUint8))
}

func main() {
	trains = make(map[uint8]Train)

	newtrain(4, "loco 2")
	fmt.Println(trains)

	newtrain(1, "loco 1")
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
func deltrain(id uint8) {
	delete(trains, id)
}

// Add a new train
func newtrain(id uint8, description string) bool {

	if _, ok := trains[id]; ok == true {
		return false
	}

	train := new(Train)
	train.description = description
	train.forward = true
	train.speed = 0

	// Store train
	trains[id] = *train

	return true
}
