package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func ComputeChecksum(file string, md5Flag, sha1Flag, sha256Flag bool) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if md5Flag {
		hash := md5.New()
		if _, err := io.Copy(hash, f); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("MD5: %x\n", hash.Sum(nil))
		f.Seek(0, 0) // Reset file pointer
	}

	if sha1Flag {
		hash := sha1.New()
		if _, err := io.Copy(hash, f); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("SHA1: %x\n", hash.Sum(nil))
		f.Seek(0, 0) // Reset file pointer
	}

	if sha256Flag {
		hash := sha256.New()
		if _, err := io.Copy(hash, f); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("SHA256: %x\n", hash.Sum(nil))
	}
}
