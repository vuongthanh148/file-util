package checksum

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestComputeChecksum(t *testing.T) {
	// Create a temporary file with known content
	content := []byte("test content")
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Expected checksums
	expectedMD5 := fmt.Sprintf("MD5: %x\n", md5.Sum(content))
	expectedSHA1 := fmt.Sprintf("SHA1: %x\n", sha1.Sum(content))
	expectedSHA256 := fmt.Sprintf("SHA256: %x\n", sha256.Sum256(content))

	tests := []struct {
		name       string
		md5Flag    bool
		sha1Flag   bool
		sha256Flag bool
		expected   string
	}{
		{"MD5", true, false, false, expectedMD5},
		{"SHA1", false, true, false, expectedSHA1},
		{"SHA256", false, false, true, expectedSHA256},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output
			var buf bytes.Buffer
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			ComputeChecksum(tmpfile.Name(), tt.md5Flag, tt.sha1Flag, tt.sha256Flag)

			w.Close()
			os.Stdout = old
			buf.ReadFrom(r)

			if buf.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, buf.String())
			}
		})
	}
}
