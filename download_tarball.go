package update_proton_ge

import (
	"io"
	"net/http"
	"os"
)

// DownloadReleaseData makes an http GET request, and returns JSON data as []byte
// that contains information about the latest Proton GE release.
func DownloadTarball(latestRelease string, tarballUrl string) error {

	file, err := os.Create("/tmp/update-proton-ge/" + latestRelease + ".tar.gz")
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(tarballUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}
