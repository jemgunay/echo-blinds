// Package main implements a test application to control two motors via GPIO pins via console commands.
package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/jemgunay/echo-blinds/motor"
	"github.com/stianeikeland/go-rpio"
)

var (
	m1 *motor.Motor
	m2 *motor.Motor
)

func main() {
	// open memory range for GPIO access
	if err := rpio.Open(); err != nil {
		log.Printf("failed to open GPIO range: %s", err)
		return
	}

	// Pi Zero pins: 16, 18, 22
	m1 = motor.New(23, 24, 25)
	// Pi Zero pins: 19, 21, 23
	m2 = motor.New(10, 9, 11)

	log.Print("> Enter command (f=forwards, b=backwards, s=stop, q=quit):")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("failed to read console input: %s", err)
			return
		}

		switch strings.ToLower(strings.TrimSpace(input)) {
		case "exit", "quit", "q":
			log.Print("> Exit")
			return

		case "forwards", "forward", "f":
			log.Print("> Forwards")
			m1.SetDirection(motor.Forwards)
			m2.SetDirection(motor.Forwards)

		case "backwards", "backward", "b":
			log.Print("> Backwards")
			m1.SetDirection(motor.Backwards)
			m2.SetDirection(motor.Backwards)

		case "stop", "s":
			log.Print("> Stop")
			m1.SetDirection(motor.None)
			m2.SetDirection(motor.None)
		}
	}
}
