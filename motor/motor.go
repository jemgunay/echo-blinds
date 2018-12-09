// Package motor implements a wrapper for the physical motor control via GPIO pins.
package motor

import (
	"context"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Direction represents the rotational direction of a motor.
type Direction string

const (
	// None indicates that the motor is powered off.
	None Direction = "none"
	// Forwards indicates that the motor is rotating in the forwards direction (dependent on wiring).
	Forwards Direction = "forwards"
	// Backwards indicates that the motor is rotating in the backwards direction (dependent on wiring).
	Backwards Direction = "backwards"
)

// Motor represents a single physical motor.
type Motor struct {
	pinA   rpio.Pin
	pinB   rpio.Pin
	pinE   rpio.Pin
	d      Direction
	cancel context.CancelFunc
}

// New initialises a new Motor and configures its output GPIO pins.
func New(pinA, pinB, pinE int) *Motor {
	m := &Motor{
		pinA: rpio.Pin(pinA),
		pinB: rpio.Pin(pinB),
		pinE: rpio.Pin(pinE),
		d:    None,
	}

	m.pinA.Mode(rpio.Output)
	m.pinB.Mode(rpio.Output)
	m.pinE.Mode(rpio.Output)
	return m
}

// SetDirection sets the rotational direction of the motor.
func (m *Motor) SetDirection(d Direction) {
	switch d {
	case None:
		m.pinA.Write(rpio.Low)
		m.pinB.Write(rpio.Low)
		m.pinE.Write(rpio.Low)
	case Forwards:
		m.pinA.Write(rpio.High)
		m.pinB.Write(rpio.Low)
		m.pinE.Write(rpio.High)
	case Backwards:
		m.pinA.Write(rpio.Low)
		m.pinB.Write(rpio.High)
		m.pinE.Write(rpio.High)
	}
	m.d = d
}

// SetDirectionWithDuration sets a motor's direction for the provided duration of time. It uses the Motor's internal
// CancelFunc to cancel any existing motor operation before initiating the new operation.
func (m *Motor) SetDirectionWithDuration(d Direction, t time.Duration) {
	m.SetDirection(d)
	if d == None {
		return
	}

	// context allows cancellation of motor operation
	var ctx = context.Background()
	ctx, m.cancel = context.WithTimeout(ctx, t)
	defer m.cancel()

	<-ctx.Done()
	m.SetDirection(None)
}

// Direction returns the rotational direction of the motor.
func (m *Motor) Direction() Direction {
	return m.d
}
