package session

import (
	privs "SWoMatic-Core/commands"
	"SWoMatic-Core/internal/io"
	"fmt"

	"go.bug.st/serial"
)

func InitiateSession(portName string, vendor string, mode serial.Mode, verbose bool) {
	switch vendor {
	case "Cisco":
		ciscoSession(portName, mode, verbose)
	default:
		// Handle unknown vendor
		return
	}
}

func ciscoSession(portName string, mode serial.Mode, verbose bool) {
	if verbose {

	}
	fmt.Println("Starting Cisco session...")
	port, err := serial.Open(portName, &mode)
	if err != nil {
		port.Close()
		return
	}


	// Dummy test
	privs.ElevatePrivilege(port, "Cisco")
	port.Write([]byte("\n"))
	response := io.ReadOutput(port)
	fmt.Println(response)
	port.Close()
}
