package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// convertToRESP converts a Redis command to RESP protocol format
func convertToRESP(command string) string {
	// Trim whitespace and split the command into parts
	parts := strings.Fields(command)

	// If no command is provided, return empty string
	if len(parts) == 0 {
		return ""
	}

	// Start building the RESP output
	var resp strings.Builder

	// Add the array length
	resp.WriteString(fmt.Sprintf("*%d\r\n", len(parts)))

	// Convert each part to a RESP bulk string
	for _, part := range parts {
		resp.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(part), part))
	}

	return resp.String()
}

func main() {
	// Check if command is provided as an argument
	if len(os.Args) > 1 {
		// Join all arguments after the program name
		command := strings.Join(os.Args[1:], " ")
		resp := convertToRESP(command)
		fmt.Println(resp)
		return
	}

	// Interactive mode
	fmt.Println("Redis RESP Converter (Press Ctrl+C to exit)")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		command := scanner.Text()
		if command == "exit" || command == "quit" {
			break
		}

		resp := convertToRESP(command)
		fmt.Println(resp)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
