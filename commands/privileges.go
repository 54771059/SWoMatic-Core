package privs

import (
	"SWoMatic-Core/internal/io"
	"bufio"
	"fmt"
	"os"
	"strings"

	"go.bug.st/serial"
)

var EntryPassword string = ""
var usedEntry bool = false
var PrivilegePassword string = ""

// detectPasswordPrompt checks if the given output contains a password prompt from a switch.
func detectPasswordPrompt(output string) bool {
	lower := strings.ToLower(output)
	// Common password prompt patterns
	prompts := []string{
		"password:",
		"enter password:",
		"please enter password:",
		"user password:",
		"admin password:",
		"passphrase:",
	}
	for _, prompt := range prompts {
		if strings.Contains(lower, prompt) {
			return true
		}
	}
	return false
}

func PasswordHandler(port serial.Port, output string, privInvoked bool) {
	// Handle password prompts and responses
	if detectPasswordPrompt(output) {
		// Respond to password prompt
		if EntryPassword != "" && !usedEntry && !privInvoked {
			// Use saved entry password
			port.Write([]byte(EntryPassword + "\n"))
			usedEntry = true
		} else if !usedEntry && !privInvoked {
			// Prompt for entry password
			fmt.Print("Enter entry password: ")
			reader := bufio.NewReader(os.Stdin)
			pass, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading entry password: %v\n", err)
				return
			}
			EntryPassword = strings.TrimSpace(pass)
			port.Write([]byte(EntryPassword + "\n"))
			usedEntry = true
		}

		if PrivilegePassword != "" {
			// Use saved privilege password
			port.Write([]byte(PrivilegePassword + "\n"))
		} else {
			// Prompt for privilege password
			fmt.Print("Enter privilege password: ")
			reader := bufio.NewReader(os.Stdin)
			pass, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading privilege password: %v\n", err)
				return
			}
			PrivilegePassword = strings.TrimSpace(pass)
			port.Write([]byte(PrivilegePassword + "\n"))
		}
	}
}

func ElevatePrivilege(port serial.Port, vendor string) {
	port.Write([]byte("\x03"))
	response := io.ReadOutput(port)
	if strings.Contains(response, "#") {
		return
	}
	switch strings.ToLower(vendor) {
	case "cisco":
		port.Write([]byte("enable\n"))
		response := io.ReadOutput(port)
		PasswordHandler(port, response, true)
	case "aruba":
		port.Write([]byte("enable\n"))
		response := io.ReadOutput(port)
		PasswordHandler(port, response, true)
	default:
		fmt.Fprintf(os.Stderr, "Unknown vendor: %s\n", vendor)
	}
}
