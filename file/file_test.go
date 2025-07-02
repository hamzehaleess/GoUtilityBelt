package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	tests := []struct {
		path     string
		expected bool
	}{
		{tmpFile.Name(), true},
		{"/nonexistent/path", false},
	}

	for _, test := range tests {
		result := Exists(test.path)
		if result != test.expected {
			t.Errorf("Exists(%q) = %v; expected %v", test.path, result, test.expected)
		}
	}
}

func TestIsDir(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	tests := []struct {
		path     string
		expected bool
	}{
		{tmpDir, true},
		{tmpFile.Name(), false},
		{"/nonexistent/path", false},
	}

	for _, test := range tests {
		result := IsDir(test.path)
		if result != test.expected {
			t.Errorf("IsDir(%q) = %v; expected %v", test.path, result, test.expected)
		}
	}
}

func TestGetExtension(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"file.txt", ".txt"},
		{"document.pdf", ".pdf"},
		{"archive.tar.gz", ".gz"},
		{"noextension", ""},
		{"/path/to/file.go", ".go"},
	}

	for _, test := range tests {
		result := GetExtension(test.path)
		if result != test.expected {
			t.Errorf("GetExtension(%q) = %q; expected %q", test.path, result, test.expected)
		}
	}
}

func TestReadWriteLines(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	lines := []string{"line1", "line2", "line3"}

	// Test WriteLines
	err = WriteLines(tmpFile.Name(), lines)
	if err != nil {
		t.Errorf("WriteLines failed: %v", err)
	}

	// Test ReadLines
	result, err := ReadLines(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadLines failed: %v", err)
	}

	if !reflect.DeepEqual(result, lines) {
		t.Errorf("ReadLines() = %v; expected %v", result, lines)
	}
}

func TestCopy(t *testing.T) {
	// Create source file
	srcFile, err := ioutil.TempFile("", "src")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(srcFile.Name())

	content := []byte("test content")
	srcFile.Write(content)
	srcFile.Close()

	// Create destination path
	dstPath := filepath.Join(os.TempDir(), "dst_test")
	defer os.Remove(dstPath)

	// Test Copy
	err = Copy(srcFile.Name(), dstPath)
	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}

	// Verify content
	dstContent, err := ioutil.ReadFile(dstPath)
	if err != nil {
		t.Errorf("Failed to read copied file: %v", err)
	}

	if !reflect.DeepEqual(dstContent, content) {
		t.Errorf("Copied content = %v; expected %v", dstContent, content)
	}
}

func TestCreateDir(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "test_create_dir", "nested", "path")
	defer os.RemoveAll(filepath.Join(os.TempDir(), "test_create_dir"))

	err := CreateDir(tmpDir)
	if err != nil {
		t.Errorf("CreateDir failed: %v", err)
	}

	if !IsDir(tmpDir) {
		t.Errorf("Directory was not created: %s", tmpDir)
	}
}
