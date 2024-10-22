package version

import "fmt"

const Version = "1.0.0"

func ShowVersion() {
	fmt.Printf("futil version %s\n", Version)
}
