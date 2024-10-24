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
	folder := filepath.Join(tmpDir, "folder")
	binaryFile := filepath.Join(tmpDir, "binaryfile.bin")

	os.WriteFile(file1, []byte("line1\nline2\nline3\n"), 0644)
	os.WriteFile(file2, []byte("line1\nline2\n"), 0644)
	os.WriteFile(file3, []byte("line1\n"), 0644)
	os.Mkdir(folder, 0755)
	os.WriteFile(binaryFile, []byte{0x00, 0x01, 0x02, 0x03}, 0644)

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
			expected:  "error: No such file '" + filepath.Join(tmpDir, "nonexistent.txt") + "'\n",
			expectErr: true,
		},
		{
			name:      "Folder check",
			file:      folder,
			expected:  "error: Expected file got directory '" + filepath.Join(tmpDir, "folder") + "'\n",
			expectErr: true,
		},
		{
			name:      "Binary file check",
			file:      binaryFile,
			expected:  "error: Cannot do linecount for binary file '" + filepath.Join(tmpDir, "binaryfile.bin") + "'\n",
			expectErr: false,
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
