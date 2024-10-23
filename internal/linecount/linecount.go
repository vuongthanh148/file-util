package linecount

import (
	"bufio"
	"fmt"
	helpers "kkcompany/helpers"
)

func CountLines(file string) {
	f, err := helpers.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the file is a binary file
	reader := bufio.NewReader(f)
	for i := 0; i < 512; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if b == 0 {
			fmt.Printf("error: Cannot do linecount for binary file '%s'\n", file)
			return
		}
	}

	// Reset the file pointer to the beginning
	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return
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
