package main

import (
	auto "SWoMatic-Core/automation"
	"SWoMatic-Core/info"
	"SWoMatic-Core/internal/constants"
	"SWoMatic-Core/internal/utils"
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	// Define flags
	listClientSerialPorts := flag.Bool("lcsp", false, "List client serial ports")
	listModes := flag.Bool("lmode", false, "List serial connection modes")
	verbose := flag.Bool("v", false, "Toggle verbose output")
	selectedMode := flag.String("mode", "cisco", "Set serial connection mode (cisco, aruba, huawei, tplink)")
	selectedClientSerialPort := flag.String("csp", "", "Set client serial port (e.g., COM3, /dev/ttyUSB0)")
	autoDetectPort := flag.Bool("auto", false, "Automatically detect connection settings")

	// Parse flags
	flag.Parse()

	// Banner
	fmt.Println("<~ SWoMatic-Core ~>")

	// Check if the -lcsp flag was passed
	if *listClientSerialPorts {
		fmt.Println(info.ListClientSerialPorts())
		os.Exit(0) // Exit after listing ports
	}

	// Check if the -lmode flag was passed
	if *listModes {
		info.ListConnectionModes()
		os.Exit(0) // Exit after listing modes
	}

	if *autoDetectPort {
		fmt.Println(auto.SwitchSweeper())
		os.Exit(0)
	}

	// Lookup selected mode & handle empty or unknown mode
	mode, ok := constants.SerialModes[*selectedMode]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown mode \"%s\"\n", *selectedMode)
		info.ListConnectionModes() // show available modes
		os.Exit(1)
	}

	// Display selected mode
	fmt.Printf("Using serial mode: %s\n", *selectedMode)
	if *verbose {
		// Mode Info
		fmt.Printf("BaudRate: %d, DataBits: %d, Parity: %s, StopBits: %s\n",
			mode.BaudRate, mode.DataBits, utils.ParityToString(mode.Parity), utils.StopBitsToString(mode.StopBits))
	}
	// Handle empty client serial port
	if *selectedClientSerialPort == "" {
		fmt.Fprintln(os.Stderr, "Error: No client serial port specified. Use the -csp flag to set it or use -auto")
		os.Exit(1)
	}

	// Validate selected client serial port
	ports := info.ListClientSerialPorts()
	found := slices.Contains(ports, *selectedClientSerialPort)
	if !found {
		fmt.Fprintf(os.Stderr, "Error: Specified client serial port \"%s\" not found.\n", *selectedClientSerialPort)
		fmt.Fprintln(os.Stderr, "Available ports:")
		for _, port := range ports {
			fmt.Fprintln(os.Stderr, "  ", port)
		}
		os.Exit(1)
	}
	fmt.Printf("Using serial port: %s\n", *selectedClientSerialPort)

}
