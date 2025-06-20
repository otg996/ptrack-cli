// Package main is the entry point for the ptrack-cli command-line tool.
// It initializes and runs the CLI application.
package main

import (
	"github.com/otg996/ptrack-cli/cmd" // Import the "cmd" package, which contains the command-line interface logic.
)

func main() {
	// Entry point of the program.
	// Calls the Run function from the "cmd" package to execute the CLI.
	cmd.Run() // Run the command-line interface.
}
