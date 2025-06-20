// Package internal provides global constant values
package internal

import "os"

// DefaultStartDirectory stores the current working directory at the time of program initialization.
// The underscore discards any potential error from os.Getwd().
var DefaultStartDirectory, _ = os.Getwd()
