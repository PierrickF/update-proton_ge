package update_proton_ge

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"os"
)

func CheckTarball(tarballFilePath string, checksumFilePath string) (err error) {
	// open tarball file
	tarball, err := os.ReadFile(tarballFilePath)
	if err != nil {
		return err
	}

	// create local checksum
	checksumArray := sha512.Sum512(tarball)
	// convert array to slice
	newChecksum := checksumArray[:]

	// compare the two checksums
	givenChecksum, err := os.ReadFile(checksumFilePath)
	if err != nil {
		return err
	}

	if !bytes.Equal(newChecksum, givenChecksum) {
		return errors.New("checksum does not match")
	}
	return nil
}
