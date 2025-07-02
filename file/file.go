// Package file provides utility functions for file and directory operations
package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir checks if a path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetSize returns the size of a file in bytes
func GetSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetExtension returns the file extension (including the dot)
func GetExtension(path string) string {
	return filepath.Ext(path)
}

// ReadLines reads all lines from a file and returns them as a slice of strings
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// WriteLines writes a slice of strings to a file, each string on a new line
func WriteLines(path string, lines []string) error {
	content := strings.Join(lines, "\n")
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// Copy copies a file from source to destination
func Copy(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, content, sourceInfo.Mode())
	return err
}

// Move moves/renames a file from source to destination
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Delete removes a file or directory
func Delete(path string) error {
	return os.RemoveAll(path)
}

// CreateDir creates a directory and all necessary parents
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// ListFiles returns a list of files in a directory
func ListFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

// ListDirs returns a list of directories in a directory
func ListDirs(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		}
	}

	return dirNames, nil
}
