package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ProcessFile checks for `{file1,file2}.ext` or `file1.ext, file2.ext`
func ProcessFile(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}

	filename := filepath.Base(filePath)
	dir := filepath.Dir(filePath)

	var validFiles []string

	// Handle `{file1,file2}.ext`
	if strings.Contains(filename, "{") && strings.Contains(filename, "}") {
		re := regexp.MustCompile(`\{([^}]*)\}(.*)`)
		matches := re.FindStringSubmatch(filename)

		if len(matches) == 3 {
			baseNames := strings.Split(matches[1], ",")
			extension := matches[2]
			for _, base := range baseNames {
				validFiles = append(validFiles, strings.TrimSpace(base)+extension)
			}
		}
	}

	// Handle `file1.ext, file2.ext`
	if strings.Contains(filename, ",") && !strings.Contains(filename, "{") {
		parts := strings.Split(filename, ",")
		for _, part := range parts {
			validFiles = append(validFiles, strings.TrimSpace(part))
		}
	}

	// Ensure at least two valid filenames exist
	if len(validFiles) < 2 {
		log.Println("Filename does not contain a valid separator")
		return
	}

	// Write content to each extracted filename
	for _, name := range validFiles {
		err := os.WriteFile(filepath.Join(dir, name), content, 0644)
		if err != nil {
			log.Printf("Error writing %s: %v", name, err)
			return
		}
	}

	// Delete the original file
	err = os.Remove(filePath)
	if err != nil {
		log.Printf("Error deleting original file: %v", err)
		return
	}

	fmt.Printf("Split %s into: %v\n", filePath, validFiles)
}
