// Package cmd provides the command-line interface functionality.
package cmd

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"

	ptrack "github.com/otg996/libptrack-go"
	"github.com/otg996/ptrack-cli/internal"
)

// Run is the entry point of the command-line interface.
// It parses command-line arguments to determine the directory to scan
// for project files.  It then scans the directory, and prints the
// project information for any projects found.  If an error occurs during
// scanning, the program will log the error to standard error and exit.
func Run() {
	// complianceTest is a flag that, when set, runs the compliance test.
	complianceTest := flag.Bool("compliance-test", false, "")

	// Parse command-line flags.
	flag.Parse()

	// If the compliance test flag is set, run the compliance test and exit.
	if *complianceTest {
		runComplianceTest()

		// Exit after the test runs.
		return
	}

	// startDir is the directory to scan for project files.
	// It defaults to internal.DefaultStartDirectory.
	// If a command-line argument is provided, it overrides the default.
	startDir := internal.DefaultStartDirectory
	if flag.NArg() > 0 {
		startDir = flag.Arg(0)
	}

	// Scan the specified directory for project files.
	projects, err := ptrack.ScanDirectory(startDir)
	if err != nil {
		// If an error occurs, log the error and exit.
		log.Fatal(err)
	}

	// Print the project information for each project found.
	for _, project := range projects {
		log.Println(project)
	}
}

// runComplianceTest runs the compliance test suite.
func runComplianceTest() {
	// referenceSuitePath is the path to the reference test suite.
	const referenceSuitePath = "./spec/reference-test-suite"

	// Prepare the test suite by creating a temporary directory.
	preparedSuitePath, err := ptrack.PrepareSuite(referenceSuitePath)
	if err != nil {
		// If an error occurs, log the error and exit.
		log.Fatalf("Failed to prepare compliance suite: %v", err)
	}
	// Must manually schedule the cleanup of the temporary directory.
	defer func() {
		// Clean up the temporary directory.
		if err := os.RemoveAll(preparedSuitePath); err != nil {
			// If an error occurs during cleanup, log the error.
			log.Printf("failed to cleanup compliance test directory: %v", err)
		}
	}()

	// Scan the prepared suite directory for projects.
	projects, err := ptrack.ScanDirectory(preparedSuitePath)
	if err != nil {
		// If an error occurs, log the error.
		log.Printf("Compliance scan failed: %v", err)

		return
	}

	// relativeProjects stores relative paths of the detected projects.
	var relativeProjects = []string{}

	// Iterate through the projects found and create relative paths.
	for _, p := range projects {
		relPath, _ := filepath.Rel(preparedSuitePath, p)
		relativeProjects = append(relativeProjects, filepath.ToSlash(relPath))
	}

	// Sort the relative project paths alphabetically.
	sort.Strings(relativeProjects)

	// Marshal the relative project paths to JSON.
	output, err := json.MarshalIndent(relativeProjects, "", "  ")
	if err != nil {
		// If an error occurs, log the error.
		log.Printf("Failed to marshal compliance output to JSON: %v", err)

		return
	}

	// Print the JSON output.
	log.Println(string(output))
}
