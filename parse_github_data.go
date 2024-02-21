package update_proton_ge

import (
	"encoding/json"
)

func ParseGithubData(data []byte) (string, error) {
	var release struct {
		ReleaseTag string `json:"tag_name"`
	}
	err := json.Unmarshal(data, &release)
	if err != nil {
		return "", err
	}

	return release.ReleaseTag, nil
}
