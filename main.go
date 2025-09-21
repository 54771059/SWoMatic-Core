package main

import (
	auto "SWoMatic-Core/automation"
	"SWoMatic-Core/info"
	"SWoMatic-Core/internal/constants"
	"SWoMatic-Core/internal/utils"
	"SWoMatic-Core/session"
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	// Define flags
	listClientSerialPorts := flag.Bool("lcsp", false, "List client serial ports")
	listModes := flag.Bool("lcmode", false, "List serial connection modes")
	verbose := flag.Bool("v", false, "Toggle verbose output")
	selectedMode := flag.String("cmode", "default", "Set serial connection mode (cisco, aruba, huawei, tplink)")
	selectedClientSerialPort := flag.String("csp", "", "Set client serial port (e.g., COM3, /dev/ttyUSB0)")
	autoDetectPort := flag.Bool("auto", false, "Automatically detect connection settings")
	selectedVendor := flag.String("vendor", "default", "Set vendor (cisco, aruba, huawei, tplink)")
	runningMode := flag.String("rmode", "", "Set running mode lbl (LineByLine)")
	filePath := flag.String("file", "", "Set config path")


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

	var switches []auto.SwitchPort
	if *autoDetectPort {
		switches = auto.SwitchSweeper()
		fmt.Println(switches)
		// TODO: Handle multi selection
		*selectedVendor = switches[0].Type
		*selectedClientSerialPort = switches[0].PortName
		*selectedMode = switches[0].ConnModeName
	}

	// Lookup selected mode & handle empty or unknown mode
	mode, ok := constants.SerialModes[*selectedMode]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown mode \"%s\"\n", *selectedMode)
		info.ListConnectionModes() // show available modes
		os.Exit(1)
	}

	// Display selected mode

	if *verbose {
		fmt.Printf("Using serial mode: %s\n", *selectedMode)
		// Mode Info
		fmt.Printf("BaudRate: %d, DataBits: %d, Parity: %s, StopBits: %s\n",
			mode.BaudRate, mode.DataBits, utils.ParityToString(mode.Parity), utils.StopBitsToString(mode.StopBits))
	}
	// Handle empty client serial port
	if *selectedClientSerialPort == "" && !*autoDetectPort {
		fmt.Fprintln(os.Stderr, "Error: No client serial port specified. Use the -csp flag to set it or use -auto")
		os.Exit(1)
	}

	if *selectedVendor == "" && !*autoDetectPort {
		fmt.Fprintln(os.Stderr, "Error: No vendor specified. Use the -vendor flag to set it or use -auto")
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
	if *verbose {
		fmt.Printf("Using serial port: %s\n", *selectedClientSerialPort)
	}

	// Validate running mode
	if *runningMode != "lbl" {
		fmt.Fprintf(os.Stderr, "Error: Unknown running mode \"%s\"\n", *runningMode)
		os.Exit(1)
	}

	if *filePath == "" && *runningMode == "lbl" {
		fmt.Fprintln(os.Stderr, "Error: No file path specified for line-by-line mode.")
		os.Exit(1)
	}

	session.InitiateSession(*selectedClientSerialPort, *selectedVendor, *mode, *verbose, *runningMode, *filePath)
}
