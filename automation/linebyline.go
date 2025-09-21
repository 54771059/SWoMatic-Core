package auto

import (
	privs "SWoMatic-Core/commands"
	"SWoMatic-Core/internal/io"
	"bufio"
	"os"
	"strings"

	"go.bug.st/serial"
)

func LineByLineReadCisco(filePath string, port serial.Port) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line
		if idx := strings.Index(line, ">"); idx != -1 {
			line = line[idx+1:]
		} else if idx := strings.Index(line, "#"); idx != -1 {
			line = line[idx+1:]
		}
		if strings.Contains(line, "enable") {
			privs.ElevatePrivilege(port, "cisco")
		} else {
			port.Write([]byte(line + "\n"))
		}
		response := io.ReadOutput(port)
		print(response)
	}
}
