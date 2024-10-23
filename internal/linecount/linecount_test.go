package linecount

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLines(t *testing.T) {
	// Create temporary test files
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")
	file3 := filepath.Join(tmpDir, "file3.txt")

	os.WriteFile(file1, []byte("line1\nline2\nline3\n"), 0644)
	os.WriteFile(file2, []byte("line1\nline2\n"), 0644)
	os.WriteFile(file3, []byte("line1\n"), 0644)

	tests := []struct {
		name      string
		file      string
		expected  string
		expectErr bool
	}{
		{
			name:      "File with 3 lines",
			file:      file1,
			expected:  "Line count: 3\n",
			expectErr: false,
		},
		{
			name:      "File with 2 lines",
			file:      file2,
			expected:  "Line count: 2\n",
			expectErr: false,
		},
		{
			name:      "File with 1 line",
			file:      file3,
			expected:  "Line count: 1\n",
			expectErr: false,
		},
		{
			name:      "Non-existent file",
			file:      filepath.Join(tmpDir, "nonexistent.txt"),
			expected:  "open " + filepath.Join(tmpDir, "nonexistent.txt") + ": no such file or directory\n",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output
			var output bytes.Buffer
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Execute the function
			CountLines(tt.file)

			// Restore stdout and read output
			w.Close()
			os.Stdout = old
			output.ReadFrom(r)

			// Verify the output
			assert.Equal(t, tt.expected, output.String())
		})
	}
}
