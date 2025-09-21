package auto

import (
	privs "SWoMatic-Core/commands"
	"SWoMatic-Core/internal/constants"
	"SWoMatic-Core/internal/io"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"go.bug.st/serial"
	// "SWoMatic-Core/info"
)

type SwitchPort struct {
	Type         string
	PortName     string
	ConnModeName string
}

func portProbeAliveCheck(portName string, connMode serial.Mode, verbose bool) (bool, string) {
	if verbose {
		fmt.Printf("Probing %s with BaudRate %d", portName, connMode.BaudRate)
		fmt.Println()
	}

	// Attempts to open connection
	port, err := serial.Open(portName, &connMode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return false, ""
	}
	// First, send carriage return and newline to wake up the switch if it's asleep
	port.Write([]byte("\r\n"))
	response := io.ReadOutput(port)
	if response != "" {
		privs.PasswordHandler(port, response, false)
	}
	// Then, send Ctrl+C to clear any leftover command or prompt
	n, err := port.Write([]byte("\x03"))
	io.ReadOutput(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		port.Close()
		return false, ""
	}

	if verbose {
		fmt.Printf("Sent %v bytes\n", n)
		fmt.Println()
	}

	if response != "" {
		if verbose {
			fmt.Println(response)
		}
		switchType := detectSwitchType(port)
		port.Close()
		return true, switchType
	}
	port.Close()
	return false, "unknown"
}

// Example placeholder function to detect switch type from response
func detectSwitchType(activePort serial.Port) string {
	activePort.Write([]byte("\x03")) // Cancel current input
	time.Sleep(200 * time.Millisecond)

	s := io.ReadOutput(activePort)

	switch {
	case regexp.MustCompile(`(?m)^<.*?>\s*$`).MatchString(s) ||
		regexp.MustCompile(`(?m)^\[.*?]\s*$`).MatchString(s):
		return "Huawei"

	case regexp.MustCompile(`(?m)^.*?>\s*$`).MatchString(s):
		return "Cisco"

	case regexp.MustCompile(`(?m)^.*?#\s*$`).MatchString(s):
		// Ambiguous: could be Cisco or Aruba — send "show version"
		activePort.Write([]byte("show version\r"))
		time.Sleep(500 * time.Millisecond)
		version := io.ReadOutput(activePort)
		if strings.Contains(version, "Cisco") {
			return "Cisco"
		}
		if strings.Contains(version, "Aruba") || strings.Contains(version, "HP") {
			return "Aruba"
		}
		return "Unknown"

	default:
		return "Unknown"
	}
}

func SwitchSweeper() []SwitchPort {
	var foundSwitches []SwitchPort
	ports, err := serial.GetPortsList()

	if len(ports) == 0 {
		fmt.Fprintf(os.Stderr, "No Ports Detected")
		return nil
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting ports list: %v\n", err)
		return nil
	}

	modes := constants.SerialModes

	for _, portName := range ports {
		var isAlive bool
		var switchType string

		// 1. Try "default" mode first
		if defaultMode, ok := modes["default"]; ok {
			isAlive, switchType = portProbeAliveCheck(portName, *defaultMode, false)
			if isAlive {
				foundSwitches = append(foundSwitches, SwitchPort{
					Type:         switchType,
					PortName:     portName,
					ConnModeName: "default",
				})
				continue // Move to next port
			}
		}

		// 2. Try remaining modes (skip "default" to avoid retrying it)
		for modeName, currMode := range modes {
			if modeName == "default" {
				continue
			}

			isAlive, switchType = portProbeAliveCheck(portName, *currMode, false)
			if isAlive {
				foundSwitches = append(foundSwitches, SwitchPort{
					Type:         switchType,
					PortName:     portName,
					ConnModeName: modeName,
				})
				break // Found a working mode, move to next port
			}
		}
	}
	return foundSwitches
}
