package linecount

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func CountLines(file string) {
	var f *os.File
	var err error

	if file == "-" {
		f = os.Stdin
	} else {
		if file == "" {
			file = "."
		}
		file = filepath.Clean(file)
		f, err = os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}

	scanner := bufio.NewScanner(f)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Line count: %d\n", lineCount)
}
