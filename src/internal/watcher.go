package internal

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher
var debounce sync.Map

// StartWatcher initializes a watcher that tracks all directories
func StartWatcher(root string) {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Add all existing directories
	err = addDirectories(root)
	if err != nil {
		log.Fatal(err)
	}

	// Listen for file system events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Create == fsnotify.Create {
				go handleNewFileOrDir(event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

// addDirectories adds all subdirectories recursively
func addDirectories(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

// handleNewFileOrDir handles new files and directories
func handleNewFileOrDir(path string) {
	info, err := os.Stat(path)
	if err != nil {
		return
	}

	// Watch new directories
	if info.IsDir() {
		_ = watcher.Add(path)
		return
	}

	// Process only files with `{}` or `,`
	if filepath.Base(path) != "" {
		debounceProcess(path)
	}
}

// debounceProcess prevents duplicate event processing
func debounceProcess(filePath string) {
	_, exists := debounce.LoadOrStore(filePath, true)
	if exists {
		return
	}

	time.Sleep(100 * time.Millisecond)
	ProcessFile(filePath)
	debounce.Delete(filePath)
}
