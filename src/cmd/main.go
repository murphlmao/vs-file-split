package main

import (
	"fmt"
	"log"
	"os"
	"vs-file-split/src/internal"
)

func main() {
	// Get the working directory
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	fmt.Printf("Monitoring directory: %s\n", rootDir)

	// Start monitoring the directory
	internal.StartWatcher(rootDir)
}
