package update_proton_ge

import (
	"net/http"
	"net/http/httputil"
)

// DownloadReleaseData makes an http GET request, and dumps the contents of the
// response as a []byte.
// It returns JSON data that contains information about the latest Proton GE
// release.
func DownloadReleaseData() (data []byte, err error) {

	resp, err := http.Get("https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest")
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
