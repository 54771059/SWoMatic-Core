package info

import (
	"SWoMatic-Core/internal/constants"
	"SWoMatic-Core/internal/utils"
	"fmt"
)

func ListConnectionModes() {
	fmt.Println("Available serial connection modes:")

	for name, mode := range constants.SerialModes {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    BaudRate: %d\n", mode.BaudRate)
		fmt.Printf("    DataBits: %d\n", mode.DataBits)
		fmt.Printf("    Parity:   %s\n", utils.ParityToString(mode.Parity))
		fmt.Printf("    StopBits: %s\n", utils.StopBitsToString(mode.StopBits))
		fmt.Println()
	}
}
