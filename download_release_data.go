package update_proton_ge

import (
	"io"
	"net/http"
)

// DownloadReleaseData makes an http GET request, and returns JSON data as []byte
// that contains information about the latest Proton GE release.
func DownloadReleaseData() (data []byte, err error) {

	resp, err := http.Get("https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest")
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
