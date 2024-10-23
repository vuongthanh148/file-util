package version

import (
	"bytes"
	"os"
	"testing"
)

func TestShowVersion(t *testing.T) {
	// Save the original os.Stdout
	originalStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	ShowVersion()

	// Close the writer and restore os.Stdout
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Check the output
	expected := "futil version " + Version + "\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
