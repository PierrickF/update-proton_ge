package update_proton_ge

import (
	"encoding/json"
	"errors"
)

// ParseReleaseData takes in JSON data downloaded from GitHub and a type of
// target to look for, and parses the data to retreive and return one of the
// following:
// - the release tag of Proton GE
// - the URL to the tarball
// - the URL to the checksum
func ParseReleaseData(data []byte, targetType string) (target string, err error) {

	type asset struct {
		Name string `json:"name"`
		Url  string `json:"browser_download_url"`
	}

	var release struct {
		ReleaseTag string  `json:"tag_name"`
		TarballUrl string  `json:"tarball_url"`
		Assets     []asset `json:"assets"`
	}

	if err = json.Unmarshal(data, &release); err != nil {
		return "", err
	}

	switch targetType {
	case "releaseTag":
		return release.ReleaseTag, nil
	case "tarballUrl":
		return release.TarballUrl, nil
	case "checksumUrl":
		return release.Assets[0].Url, nil
	default:
		return "", errors.New("invalid target type")
	}
}
