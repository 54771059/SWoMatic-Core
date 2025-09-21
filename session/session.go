package session

import (
	auto "SWoMatic-Core/automation"
	"fmt"

	"go.bug.st/serial"
)

func InitiateSession(portName string, vendor string, mode serial.Mode, verbose bool, runningMode string, filePath string) {
	switch vendor {
	case "Cisco":
		ciscoSession(portName, mode, verbose, runningMode, filePath)
	default:
		// Handle unknown vendor
		return
	}
}

func ciscoSession(portName string, mode serial.Mode, verbose bool, runningMode string, filePath string) {
	if verbose {

	}
	fmt.Println("Starting Cisco session...")
	port, err := serial.Open(portName, &mode)
	if err != nil {
		port.Close()
		return
	}

	// Dummy test
	if runningMode == "lbl" {
		auto.LineByLineReadCisco(filePath, port)
	}

	port.Close()
}
