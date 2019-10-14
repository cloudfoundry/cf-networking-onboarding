package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Build a CSV for import to Pivotal Tracker by uncommenting the track you'd like to output

func main() {
	modules := []string{
		"introduction.prolific.md",
		"routes.prolific.md",
		"route-integrity.prolific.md",
		"c2c.prolific.md",
		"c2c-service-discovery.prolific.md",
		"iptables-primer.prolific.md",
		"asgs.prolific.md",
		"route-registrar.prolific.md",
		"bosh-dns.prolific.md",
		"tcp-routes.prolific.md",
		"route-services.prolific.md",
		"interrupts.prolific.md",
		"write-your-own-story.prolific.md",
	}

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filepath.Join(workingDir, "networking-program-onboarding-tracker.csv"))
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	for i, module := range modules {
		moduleFilePath := filepath.Join(workingDir, module)
		cmd := exec.Command("prolific", moduleFilePath)
		cmd.Stderr = os.Stderr
		csvContent, err := cmd.Output()
		if err != nil {
			log.Fatalf("Failed to run prolific: %s", err)
		}

		if i != 0 {
			csvContent = bytes.TrimLeft(csvContent, "Title, Type, Description, Labels,Task\n")
		}

		_, err = outputFile.Write(csvContent)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Generating CSV with selected modules: %s", modules)
}
