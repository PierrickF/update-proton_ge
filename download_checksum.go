package update_proton_ge

import (
	"io"
	"net/http"
	"os"
)

// DownloadChecksum creates a file and populates it with the contents of an
// HTTP request.
// The result is a sha512 file for the latest Proton GE release.
func DownloadChecksum(latestRelease string, checksumUrl string) error {

	file, err := os.Create("/tmp/update-proton-ge/" + latestRelease + ".sha512sum")
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(checksumUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}
