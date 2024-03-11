package update_proton_ge

import (
	"bytes"
	"crypto/sha512"
	"os"
)

func CheckTarball(tarballFilePath string, checksumFilePath string) (isCorrect bool, err error) {
	// open tarball file
	tarball, err := os.ReadFile(tarballFilePath)
	if err != nil {
		return false, err
	}

	// create local checksum
	checksumArray := sha512.Sum512(tarball)
	// convert array to slice
	newChecksum := checksumArray[:]

	// compare the two checksums
	givenChecksum, err := os.ReadFile(checksumFilePath)
	if err != nil {
		return false, err
	}

	return bytes.Equal(newChecksum, givenChecksum), nil
}
