package main

import (
	"SWoMatic-Core/info"
	"SWoMatic-Core/internal/constants"
	"SWoMatic-Core/internal/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	listClientSerialPorts := flag.Bool("lcsp", false, "List client serial ports")
	listModes := flag.Bool("lmode", false, "List serial connection modes")
	verbose := flag.Bool("v", false, "Toggle verbose output")
	selectedMode := flag.String("mode", "cisco", "Set serial connection mode (cisco, aruba, huawei, tplink)")

	// Parse flags
	flag.Parse()

	fmt.Println("SWoMatic-Core ~>")

	// Check if the -lcsp flag was passed
	if *listClientSerialPorts {
		info.ListClientSerialPorts()
		os.Exit(0) // Exit after listing ports
	}

	if *listModes {
		info.ListConnectionModes()
		os.Exit(0) // Exit after listing modes
	}

	// Lookup selected mode
	mode, ok := constants.SerialModes[*selectedMode]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown mode \"%s\"\n", *selectedMode)
		info.ListConnectionModes() // show available modes
		os.Exit(1)
	}

	if *verbose {
		// Mode Info
		fmt.Printf("Using serial mode: %s\n", *selectedMode)
		fmt.Printf("BaudRate: %d, DataBits: %d, Parity: %s, StopBits: %s\n",
			mode.BaudRate, mode.DataBits, utils.ParityToString(mode.Parity), utils.StopBitsToString(mode.StopBits))
	}

}
