package train

import (
	"math"
)

type Train struct {
	id          uint8
	Description string
	Forward     bool
	Speed       uint8
}

func (t *Train) DirectionChange() {
	t.Forward = !t.Forward
}

func (t *Train) Stop() {
	t.SpeedChange(0)
}

func (t *Train) SpeedChange(speed uint8) {
	// Actual stop code
	// ...
	// Change speed
	t.Speed = speed
}

func (t *Train) RelativeSpeedChange(speed int) uint8 {
	newspeed := NormalizeSpeed(int(t.Speed) + speed)

	t.SpeedChange(newspeed)
	return newspeed
}

func NormalizeSpeed(speed int) uint8 {
	return uint8(math.Min((math.Max(float64(speed), 0)), math.MaxUint8))
}
