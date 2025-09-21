package privs

import (
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

func PasswordHandler(port serial.Port, output string) {
	// Handle password prompts and responses
	if detectPasswordPrompt(output) {
		// Respond to password prompt
		if EntryPassword != "" && !usedEntry {
			// Use saved entry password
			port.Write([]byte(EntryPassword + "\n"))
			usedEntry = true
		} else {
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