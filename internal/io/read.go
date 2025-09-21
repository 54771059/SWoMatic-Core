package io

import (
	"fmt"
	"os"
	"time"

	"go.bug.st/serial"
)

func ReadOutput(port serial.Port) string {
	buff := make([]byte, 100)
	response := ""
	for {
		port.SetReadTimeout(time.Millisecond * 500)
		n, err := port.Read(buff)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			break
		}
		if n == 0 {
			break
		}
		response += string(buff[:n])
	}

	return response
}
