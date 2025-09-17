package utils

import (
	"go.bug.st/serial"
)

// helper function to format parity
func ParityToString(p serial.Parity) string {
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
func StopBitsToString(s serial.StopBits) string {
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
