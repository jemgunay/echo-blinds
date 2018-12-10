// Package blinds implements the physical blinds opening and closing controls for an Amazon Echo skill.
package blinds

import (
	"fmt"
	"log"
	"time"

	"github.com/jemgunay/echo-blinds/motor"
	"github.com/stianeikeland/go-rpio"
)

// State is used to track the current and target blinds state.
type State string

const (
	// Open state implies the blinds are fully opened.
	Open State = "open"
	// Close state implies the blinds are fully closed.
	Close State = "closed"
	// Stop state implies the blinds were stopped mid open or close.
	Stop State = "cancel"
)

var (
	// rotate motor rotates the blinds (required before drawing the blinds)
	rotateMotor *motor.Motor
	// draw motor draws the blinds linearly
	drawMotor *motor.Motor

	// state change request channel
	stateCh = make(chan State)
	// current blinds state is initialised to Stop as the starting blinds state is unknown
	currentState = Stop
)

// Init sets up GPIO pin access, initialises both motors and starts the update poller.
func Init() error {
	// open memory range for GPIO access
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open GPIO range: %s", err)
	}

	// Pi Zero pins: 16, 18, 22
	rotateMotor = motor.New(23, 24, 25)
	// Pi Zero pins: 19, 21, 23
	drawMotor = motor.New(10, 9, 11)

	// poll for update requests
	go updatePoller()
	return nil
}

// Update sets the target blinds state. It serialises requests to update the blinds state, preventing accidental
// concurrent access. The response text is spoken by Alexa.
func Update(state State) (response string) {
	// check if blinds is already in the provided update state
	if currentState == state && state != Stop {
		switch state {
		case Open:
			return "The blinds are already open."
		case Close:
			return "The blinds are already closed."
		}
	}

	currentState = state
	stateCh <- state

	// new state successfully set
	switch state {
	case Open:
		return "Opening the blinds."
	case Close:
		return "Closing the blinds."
	default:
		return "Stopping the blinds."
	}
}

// polls for new updates
func updatePoller() {
	for state := range stateCh {
		// cancels any existing motor operation
		rotateMotor.SetDirectionWithDuration(motor.None, 0)
		drawMotor.SetDirectionWithDuration(motor.None, 0)
		time.Sleep(time.Second)

		switch state {
		case Open:
			log.Print("Opening blinds...")
			go rotateMotor.SetDirectionWithDuration(motor.Forwards, time.Second*2)
			go drawMotor.SetDirectionWithDuration(motor.Forwards, time.Second*8)
			log.Print("Blinds opened.")

		case Close:
			log.Print("Closing blinds...")
			go rotateMotor.SetDirectionWithDuration(motor.Backwards, time.Second*2)
			go drawMotor.SetDirectionWithDuration(motor.Backwards, time.Second*8)
			log.Print("Blinds closed.")

		case Stop:
			log.Print("Stopped blinds.")

		default:
			log.Printf("Unsupported blinds command: %s", state)
		}
	}
}

// Shutdown closes the state update channel and closes the GPIO pins.
func Shutdown() {
	close(stateCh)
	// clean up GPIO pins
	if err := rpio.Close(); err != nil {
		log.Printf("failed to close GPIO range: %s", err)
	}
}
