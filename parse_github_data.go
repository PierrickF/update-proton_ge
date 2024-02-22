package update_proton_ge

import (
	"encoding/json"
)

// ParseGithubData takes in JSON data downloaded from GitHub,
// and parses it to retreive the release name of Proton GE,
// which is then returned.
func ParseGithubData(data []byte) (name string, err error) {
	var release struct {
		ReleaseTag string `json:"tag_name"`
	}
	err = json.Unmarshal(data, &release)
	if err != nil {
		return "", err
	}

	return release.ReleaseTag, nil
}
