package update_proton_ge

import (
	"io"
	"net/http"
	"os"
)

// DownloadTarball creates a file and populates it with the contents of an
// HTTP request.
// The result is an archive file of the latest Proton GE release.
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
