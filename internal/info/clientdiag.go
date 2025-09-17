package info

import (
	"fmt"

	"go.bug.st/serial"
)

func ListClientSerialPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println("Error listing serial ports:", err)
		return nil
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found")
		return nil
	}
	fmt.Println("Available serial ports:")
	var result []string
	for _, port := range ports {
		fmt.Println(" -", port)
		result = append(result, port)
	}
	return result
}