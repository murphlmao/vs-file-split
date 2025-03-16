package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test `{file1,file2}.ext` pattern
func TestSplitCurlyBraceFiles(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "{test1,test2}.txt")
	os.WriteFile(filePath, []byte("Hello World"), 0644)

	ProcessFile(filePath)

	assert.FileExists(t, filepath.Join(tmpDir, "test1.txt"))
	assert.FileExists(t, filepath.Join(tmpDir, "test2.txt"))

	_, err := os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}

// Test `file1.ext, file2.ext` pattern
func TestSplitCommaFiles(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fileA.txt, fileB.txt")
	os.WriteFile(filePath, []byte("Sample Text"), 0644)

	ProcessFile(filePath)

	assert.FileExists(t, filepath.Join(tmpDir, "fileA.txt"))
	assert.FileExists(t, filepath.Join(tmpDir, "fileB.txt"))

	_, err := os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}
