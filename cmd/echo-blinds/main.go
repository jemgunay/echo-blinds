// Package main implements the HTTP server which handles requests for opening and closing window blinds as part of an
// an amazon echo skill.
package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/jemgunay/echo-blinds"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

// skill ID is found on the alexa developer console
const skillID = "enter skill ID here"

func main() {
	// parse flags
	port := flag.Uint64("port", 3000, "the port for the HTTP server to listen on")

	// init blinds control
	if err := blinds.Init(); err != nil {
		log.Printf("failed to initialise blinds: %s", err)
		return
	}
	defer blinds.Shutdown()

	// define handler routing
	echoApps := map[string]interface{}{
		"/echo/blinds": alexa.EchoApplication{
			AppID:    skillID,
			OnIntent: echoIntentHandler,
			OnLaunch: echoIntentHandler,
		},
	}

	// init Alexa interface
	alexa.Run(echoApps, strconv.FormatUint(*port, 10))
}

// routes blinds requests and handled writing Alexa speech responses
func echoIntentHandler(r *alexa.EchoRequest, w *alexa.EchoResponse) {
	switch r.GetIntentName() {
	case "OpenBlindsIntent":
		w.OutputSpeech(blinds.Update(blinds.Open))
	case "CloseBlindsIntent":
		w.OutputSpeech(blinds.Update(blinds.Close))
	case "StopBlindsIntent":
		w.OutputSpeech(blinds.Update(blinds.Stop))
	default:
		w.OutputSpeech("That's not a supported blinds command.")
	}
}
