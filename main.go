package main

import (
	"SWoMatic-Core/internal/info"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the flag
	listClientSerialPorts := flag.Bool("lcsp", false, "List client serial ports")
	listConnetionModes := flag.Bool("lmod", false, "List client serial ports")

	// Parse flags
	flag.Parse()

	fmt.Println("SWoMatic-Core ~>")

	// Check if the -lcsp flag was passed
	if *listClientSerialPorts {
		info.ListClientSerialPorts()
		os.Exit(0) // Exit after listing ports
	}

	if *listConnetionModes {
		info.ListConnectionModes()
		os.Exit(0) // Exit after listing ports
	}
}
