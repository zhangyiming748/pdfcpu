package util

import (
	"os"
	"path/filepath"
	"strings"
)

// GetPDFFiles returns a list of all PDF files in the given directory and its subdirectories,
// ignoring case sensitivity of the file extension.
func GetPDFFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file has .pdf extension (case insensitive)
		if strings.ToLower(filepath.Ext(path)) == ".pdf" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
