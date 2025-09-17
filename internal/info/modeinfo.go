package info

import (
	"SWoMatic-Core/internal/constants"
	"fmt"

	"go.bug.st/serial"
)

func ListConnectionModes() {
	fmt.Println("Available serial connection modes:")

	for name, mode := range constants.SerialModes {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    BaudRate: %d\n", mode.BaudRate)
		fmt.Printf("    DataBits: %d\n", mode.DataBits)
		fmt.Printf("    Parity:   %s\n", parityToString(mode.Parity))
		fmt.Printf("    StopBits: %s\n", stopBitsToString(mode.StopBits))
		fmt.Println()
	}
}

// helper function to format parity
func parityToString(p serial.Parity) string {
	switch p {
	case serial.NoParity:
		return "None"
	case serial.OddParity:
		return "Odd"
	case serial.EvenParity:
		return "Even"
	default:
		return "Unknown"
	}
}

// helper function to format stop bits
func stopBitsToString(s serial.StopBits) string {
	switch s {
	case serial.OneStopBit:
		return "1"
	case serial.OnePointFiveStopBits:
		return "1.5"
	case serial.TwoStopBits:
		return "2"
	default:
		return "Unknown"
	}
}
