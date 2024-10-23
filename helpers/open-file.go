package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

// openFile handles the file opening and validation logic
func OpenFile(file string) (*os.File, error) {
	if file == "-" {
		return os.Stdin, nil
	}

	if file == "" {
		file = "."
	}
	file = filepath.Clean(file)

	info, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error: No such file '%s'", file)
		}
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("error: Expected file got directory '%s'", file)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return f, nil
}
