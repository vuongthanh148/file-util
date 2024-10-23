package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeChecksum(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		md5Flag    bool
		sha1Flag   bool
		sha256Flag bool
		expected   string
	}{
		{
			name:       "MD5 checksum",
			content:    "hello",
			md5Flag:    true,
			sha1Flag:   false,
			sha256Flag: false,
			expected:   fmt.Sprintf("MD5: %x\n", md5.Sum([]byte("hello"))),
		},
		{
			name:       "SHA1 checksum",
			content:    "hello",
			md5Flag:    false,
			sha1Flag:   true,
			sha256Flag: false,
			expected:   fmt.Sprintf("SHA1: %x\n", sha1.Sum([]byte("hello"))),
		},
		{
			name:       "SHA256 checksum",
			content:    "hello",
			md5Flag:    false,
			sha1Flag:   false,
			sha256Flag: true,
			expected:   fmt.Sprintf("SHA256: %x\n", sha256.Sum256([]byte("hello"))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the test content
			tmpFile, err := ioutil.TempFile("", "testfile")
			assert.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.WriteString(tt.content)
			assert.NoError(t, err)
			tmpFile.Close()

			// Capture the output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Execute the function
			ComputeChecksum(tmpFile.Name(), tt.md5Flag, tt.sha1Flag, tt.sha256Flag)

			// Restore stdout and read output
			w.Close()
			os.Stdout = old
			var output []byte
			output, err = ioutil.ReadAll(r)
			assert.NoError(t, err)

			// Verify the output
			assert.Equal(t, tt.expected, string(output))
		})
	}
}
