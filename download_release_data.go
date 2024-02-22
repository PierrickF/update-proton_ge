package update_proton_ge

import (
	"fmt"
	"os/exec"
)

// GetLatestVersionName downloads a JSON file corresponding to the latest
// release of Proton GE, and feeds it to ParseGithubData.
// It returns the latest version name of Proton GE.
func GetLatestVersionName() (name string, err error) {

	curlUrl := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"
	curl := exec.Command("curl", "-s", curlUrl)

	curlOutput, err := curl.Output()
	if err != nil {
		return "", fmt.Errorf("curl error: %w", err)
	}

	return ParseGithubData(curlOutput)
}
