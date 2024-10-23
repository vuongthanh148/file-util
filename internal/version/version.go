package version

import "fmt"

var Version = ""

func ShowVersion() {
	fmt.Printf("futil version %s\n", Version)
}
