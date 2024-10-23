package version

import "fmt"

var Version = "v1.0.11"

func ShowVersion() {
	fmt.Printf("futil %s\n", Version)
}
